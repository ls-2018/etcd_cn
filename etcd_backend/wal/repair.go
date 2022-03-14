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

package wal

import (
	"github.com/ls-2018/etcd_cn/client_sdk/pkg/fileutil"
	"github.com/ls-2018/etcd_cn/etcd_backend/wal/walpb"
	"go.uber.org/zap"
	"io"
	"os"
	"path/filepath"
)

// Repair tries to repair ErrUnexpectedEOF in the
// last wal file by truncating.
func Repair(lg *zap.Logger, dirpath string) bool {
	if lg == nil {
		lg = zap.NewNop()
	}
	f, err := openLast(lg, dirpath)
	if err != nil {
		return false
	}
	defer f.Close()

	lg.Info("repairing", zap.String("path", f.Name()))

	rec := &walpb.Record{}
	decoder := newDecoder(f)
	for {
		lastOffset := decoder.lastOffset()
		err := decoder.decode(rec)
		switch err {
		case nil:
			// update crc of the decoder when necessary
			switch rec.Type {
			case crcType:
				crc := decoder.crc.Sum32()
				// current crc of decoder must match the crc of the record.
				// do no need to match 0 crc, since the decoder is a new one at this case.
				if crc != 0 && rec.Validate(crc) != nil {
					return false
				}
				decoder.updateCRC(rec.Crc)
			}
			continue

		case io.EOF:
			lg.Info("repaired", zap.String("path", f.Name()), zap.Error(io.EOF))
			return true

		case io.ErrUnexpectedEOF:
			bf, bferr := os.Create(f.Name() + ".broken")
			if bferr != nil {
				lg.Warn("failed to create backup file", zap.String("path", f.Name()+".broken"), zap.Error(bferr))
				return false
			}
			defer bf.Close()

			if _, err = f.Seek(0, io.SeekStart); err != nil {
				lg.Warn("failed to read file", zap.String("path", f.Name()), zap.Error(err))
				return false
			}

			if _, err = io.Copy(bf, f); err != nil {
				lg.Warn("failed to copy", zap.String("from", f.Name()+".broken"), zap.String("to", f.Name()), zap.Error(err))
				return false
			}

			if err = f.Truncate(lastOffset); err != nil {
				lg.Warn("failed to truncate", zap.String("path", f.Name()), zap.Error(err))
				return false
			}

			if err = fileutil.Fsync(f.File); err != nil {
				lg.Warn("failed to fsync", zap.String("path", f.Name()), zap.Error(err))
				return false
			}
			lg.Info("repaired", zap.String("path", f.Name()), zap.Error(io.ErrUnexpectedEOF))
			return true

		default:
			lg.Warn("failed to repair", zap.String("path", f.Name()), zap.Error(err))
			return false
		}
	}
}

// openLast opens the last wal file for read and write.
func openLast(lg *zap.Logger, dirpath string) (*fileutil.LockedFile, error) {
	names, err := readWALNames(lg, dirpath)
	if err != nil {
		return nil, err
	}
	last := filepath.Join(dirpath, names[len(names)-1])
	return fileutil.LockFile(last, os.O_RDWR, fileutil.PrivateFileMode)
}
