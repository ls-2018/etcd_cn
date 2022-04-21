// Copyright 2015 The etcd Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mvcc

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/btree"
	"go.uber.org/zap"
)

var ErrRevisionNotFound = errors.New("mvcc: 修订版本没有找到")

// keyIndex
// key的删除操作将在末尾追加删除版本
// 当前代,并创建一个新的空代.
type keyIndex struct {
	key         string       // key
	modified    revision     // 一个key 最新修改的revision .
	generations []generation // 每次新建都会创建一个,删除然后新建也会生成一个
}

// generation  包含一个key的多个版本.
type generation struct {
	versionCount int64      // 记录对当前key 有几个版本
	created      revision   // 第一次创建时的索引信息
	revs         []revision // 当值存在以后,对该值的修改记录
}

type revision struct {
	main int64 // 一个全局递增的主版本号,随put/txn/delete事务递增,一个事务内的key main版本号是一致的
	sub  int64 // 一个事务内的子版本号,从0开始随事务内put/delete操作递增
}

func (ki *keyIndex) restore(lg *zap.Logger, created, modified revision, ver int64) {
	if len(ki.generations) != 0 {
		lg.Panic(
			"'restore' got an unexpected non-empty generations",
			zap.Int("generations-size", len(ki.generations)),
		)
	}
	ki.modified = modified
	g := generation{created: created, versionCount: ver, revs: []revision{modified}}
	ki.generations = append(ki.generations, g)
}

// tombstone puts a revision, pointing to a tombstone, to the keyIndex.
// It also creates a new empty generation in the keyIndex.
// It returns ErrRevisionNotFound when tombstone on an empty generation.
func (ki *keyIndex) tombstone(lg *zap.Logger, main int64, sub int64) error {
	// 当然如果 keyIndex 中的最大版本号被打了删除标记 (tombstone)， 就会从 treeIndex 中删除这个 keyIndex，否则会出现内存泄露。
	if ki.isEmpty() {
		lg.Panic(
			"'tombstone' got an unexpected empty keyIndex",
			zap.String("key", string(ki.key)),
		)
	}
	if ki.generations[len(ki.generations)-1].isEmpty() {
		return ErrRevisionNotFound
	}
	ki.put(lg, main, sub)
	ki.generations = append(ki.generations, generation{})
	return nil
}

// since returns revisions since the given rev. Only the revision with the
// largest sub revision will be returned if multiple revisions have the same
// main revision.
func (ki *keyIndex) since(lg *zap.Logger, rev int64) []revision {
	if ki.isEmpty() {
		lg.Panic("'since' 得到一个意外的空keyIndex", zap.String("key", ki.key))
	}
	since := revision{rev, 0}
	var gi int
	// find the generations to start checking
	for gi = len(ki.generations) - 1; gi > 0; gi-- {
		g := ki.generations[gi]
		if g.isEmpty() {
			continue
		}
		if since.GreaterThan(g.created) {
			break
		}
	}

	var revs []revision
	var last int64
	for ; gi < len(ki.generations); gi++ {
		for _, r := range ki.generations[gi].revs {
			if since.GreaterThan(r) {
				continue
			}
			if r.main == last {
				// replace the revision with a new one that has higher sub value,
				// because the original one should not be seen by external
				revs[len(revs)-1] = r
				continue
			}
			revs = append(revs, r)
			last = r.main
		}
	}
	return revs
}

// compact compacts a keyIndex by removing the versions with smaller or equal
// revision than the given atRev except the largest one (If the largest one is
// a tombstone, it will not be kept).
// If a generation becomes empty during compaction, it will be removed.
func (ki *keyIndex) compact(lg *zap.Logger, atRev int64, available map[revision]struct{}) {
	if ki.isEmpty() {
		lg.Panic(
			"'compact' got an unexpected empty keyIndex",
			zap.String("key", string(ki.key)),
		)
	}

	genIdx, revIndex := ki.doCompact(atRev, available)

	g := &ki.generations[genIdx]
	if !g.isEmpty() {
		// remove the previous contents.
		if revIndex != -1 {
			g.revs = g.revs[revIndex:]
		}
		// remove any tombstone
		if len(g.revs) == 1 && genIdx != len(ki.generations)-1 {
			delete(available, g.revs[0])
			genIdx++
		}
	}

	// remove the previous generations.
	ki.generations = ki.generations[genIdx:]
}

// keep finds the revision to be kept if compact is called at given atRev.
func (ki *keyIndex) keep(atRev int64, available map[revision]struct{}) {
	if ki.isEmpty() {
		return
	}

	genIdx, revIndex := ki.doCompact(atRev, available)
	g := &ki.generations[genIdx]
	if !g.isEmpty() {
		// remove any tombstone
		if revIndex == len(g.revs)-1 && genIdx != len(ki.generations)-1 {
			delete(available, g.revs[revIndex])
		}
	}
}

func (ki *keyIndex) doCompact(atRev int64, available map[revision]struct{}) (genIdx int, revIndex int) {
	// walk until reaching the first revision smaller or equal to "atRev",
	// and add the revision to the available map
	f := func(rev revision) bool {
		if rev.main <= atRev {
			available[rev] = struct{}{}
			return false
		}
		return true
	}

	genIdx, g := 0, &ki.generations[0]
	// find first generation includes atRev or created after atRev
	for genIdx < len(ki.generations)-1 {
		if tomb := g.revs[len(g.revs)-1].main; tomb > atRev {
			break
		}
		genIdx++
		g = &ki.generations[genIdx]
	}

	revIndex = g.walk(f)

	return genIdx, revIndex
}

// --------------------------------------------- OVER ---------------------------------------------------------------

// get 获取满足给定atRev的键的修改、创建的revision和版本.Rev必须大于或等于给定的atRev.
func (ki *keyIndex) get(lg *zap.Logger, atRev int64) (modified, created revision, ver int64, err error) {
	if ki.isEmpty() { // 判断有没有修订版本
		lg.Panic("'get'得到一个意外的空keyIndex", zap.String("key", ki.key))
	}
	g := ki.findGeneration(atRev)
	if g.isEmpty() {
		return revision{}, revision{}, 0, ErrRevisionNotFound
	}

	n := g.walk(func(rev revision) bool { return rev.main > atRev }) // 返回第一个小于等于 该修订版本的最新的索引
	if n != -1 {
		return g.revs[n], g.created, g.versionCount - int64(len(g.revs)-n-1), nil
	}

	return revision{}, revision{}, 0, ErrRevisionNotFound
}

func (ki *keyIndex) Less(b btree.Item) bool {
	return strings.Compare(ki.key, b.(*keyIndex).key) == -1
}

func (ki *keyIndex) equal(b *keyIndex) bool {
	if !strings.EqualFold(ki.key, b.key) {
		return false
	}
	if ki.modified != b.modified {
		return false
	}
	if len(ki.generations) != len(b.generations) {
		return false
	}
	for i := range ki.generations {
		ag, bg := ki.generations[i], b.generations[i]
		if !ag.equal(bg) {
			return false
		}
	}
	return true
}

func (ki *keyIndex) String() string {
	var s string
	for _, g := range ki.generations {
		s += g.String()
	}
	return s
}

func (g *generation) isEmpty() bool { return g == nil || len(g.revs) == 0 }

// 遍历返回符合条件的索引,倒序遍历
func (g *generation) walk(f func(rev revision) bool) int {
	l := len(g.revs)
	for i := range g.revs {
		ok := f(g.revs[l-i-1])
		if !ok {
			return l - i - 1
		}
	}
	return -1
}

func (g generation) equal(b generation) bool {
	if g.versionCount != b.versionCount {
		return false
	}
	if len(g.revs) != len(b.revs) {
		return false
	}

	for i := range g.revs {
		ar, br := g.revs[i], b.revs[i]
		if ar != br {
			return false
		}
	}
	return true
}

func (g *generation) String() string {
	return fmt.Sprintf("g: 创建[%d] 版本数[%d], 修订记录 %#v\n", g.created, g.versionCount, g.revs)
}

// OK
func (ki *keyIndex) isEmpty() bool {
	// 只有一个历史版本,且
	return len(ki.generations) == 1 && ki.generations[0].isEmpty()
}

// findGeneration 找到给定rev所属的keyIndex的生成.如果给定的rev在两代之间,这意味着在给定的rev上键不存在,它将返回nil.
// 如果修订版本是接下来要写的,就返回当前代
func (ki *keyIndex) findGeneration(rev int64) *generation {
	lastg := len(ki.generations) - 1
	cg := lastg
	// 倒着查找
	for cg >= 0 {
		if len(ki.generations[cg].revs) == 0 {
			cg--
			continue
		}
		g := ki.generations[cg]
		if cg != lastg {
			// 每次生成的key的最新修订版本
			// 不是最新的一组,但最大的都比 rev
			if tomb := g.revs[len(g.revs)-1].main; tomb <= rev {
				// tomb应该是删除的代数
				return nil
			}
		}
		// 0    rev     last
		if g.revs[0].main <= rev {
			// 找到对应修订版本 所属的gen版本
			return &ki.generations[cg]
		}
		cg--
	}
	return nil
}

// put 将一个修订放到keyIndex中.
func (ki *keyIndex) put(lg *zap.Logger, main int64, sub int64) {
	rev := revision{main: main, sub: sub}
	if !rev.GreaterThan(ki.modified) {
		lg.Panic(
			"'put'有一个意想不到的小修改",
			zap.Int64("given-revision-main", rev.main),
			zap.Int64("given-revision-sub", rev.sub),
			zap.Int64("modified-revision-main", ki.modified.main),
			zap.Int64("modified-revision-sub", ki.modified.sub),
		)
	}
	if len(ki.generations) == 0 {
		ki.generations = append(ki.generations, generation{})
	}
	g := &ki.generations[len(ki.generations)-1]
	if len(g.revs) == 0 { // create a new key
		g.created = rev
	}
	g.revs = append(g.revs, rev)
	g.versionCount++
	ki.modified = rev
}
