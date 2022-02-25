# etcd

[![Go Report Card](https://goreportcard.com/badge/github.com/etcd-io/etcd?style=flat-square)](https://goreportcard.com/report/github.com/etcd-io/etcd)
[![Coverage](https://codecov.io/gh/etcd-io/etcd/branch/master/graph/badge.svg)](https://codecov.io/gh/etcd-io/etcd)
[![Tests](https://github.com/etcd-io/etcd/actions/workflows/tests.yaml/badge.svg)](https://github.com/etcd-io/etcd/actions/workflows/tests.yaml)
[![asset-transparency](https://github.com/etcd-io/etcd/actions/workflows/asset-transparency.yaml/badge.svg)](https://github.com/etcd-io/etcd/actions/workflows/asset-transparency.yaml)
[![codeql-analysis](https://github.com/etcd-io/etcd/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/etcd-io/etcd/actions/workflows/codeql-analysis.yml)
[![self-hosted-linux-arm64-graviton2-tests](https://github.com/etcd-io/etcd/actions/workflows/self-hosted-linux-arm64-graviton2-tests.yml/badge.svg)](https://github.com/etcd-io/etcd/actions/workflows/self-hosted-linux-arm64-graviton2-tests.yml)
[![Docs](https://img.shields.io/badge/docs-latest-green.svg)](https://etcd.io/docs)
[![Godoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/etcd-io/etcd)
[![Releases](https://img.shields.io/github/release/etcd-io/etcd/all.svg?style=flat-square)](https://github.com/etcd-io/etcd/releases)
[![LICENSE](https://img.shields.io/github/license/etcd-io/etcd.svg?style=flat-square)](https://github.com/etcd-io/etcd/blob/main/LICENSE)

**Note**: The `main` branch may be in an *unstable or even broken state* during development. For stable versions,
see [releases][github-release].

![etcd Logo](logos/etcd-horizontal-color.svg)

etcd is a distributed reliable key-value store for the most critical data of a distributed system, with a focus on
being:

* *Simple*: well-defined, user-facing API (gRPC)
* *Secure*: automatic TLS with optional client cert authentication
* *Fast*: benchmarked 10,000 writes/sec
* *Reliable*: properly distributed using Raft

etcd is written in Go and uses the [Raft][] consensus algorithm to manage a highly-available replicated log.

etcd is used [in production by many companies](./ADOPTERS.md), and the development team stands behind it in critical
deployment scenarios, where etcd is frequently teamed with applications such as [Kubernetes][k8s], [locksmith][]
, [vulcand][], [Doorman][], and many others. Reliability is further ensured by [**rigorous
testing**](https://github.com/etcd-io/etcd/tree/main/tests/functional).

See [etcdctl][etcdctl] for a simple command line client.

[raft]: https://raft.github.io/

[k8s]: http://kubernetes.io/

[doorman]: https://github.com/youtube/doorman

[locksmith]: https://github.com/coreos/locksmith

[vulcand]: https://github.com/vulcand/vulcand

[etcdctl]: https://github.com/etcd-io/etcd/tree/main/etcdctl

```
自己看源码 用
```