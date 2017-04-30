#!/bin/bash

set -e
set -x

mkdir test

alias tbm='the-busy-man'

tbm -h
tbm -help
tbm -v
tbm -version
tbm -l
tbm -h git

export VERBOSE=y

rm -fr test && mkdir test
tbm -w test git
(ls -alh test | grep .git) || exit 1

rm -fr test && mkdir test
tbm -w test git license:mit
(ls -alh test | grep .git) || exit 1
(ls -alh test | grep LICENSE) || exit 1

rm -fr test && mkdir test
tbm -w test git license:mit git:init+commit changelog
(ls -alh test | grep .git) || exit 1
(ls -alh test | grep LICENSE) || exit 1
(ls -alh test | grep change.log) || exit 1

rm -fr test && mkdir test
tbm -w test git license:mit golang
(ls -alh test | grep .git) || exit 1
(ls -alh test | grep main.go) || exit 1

rm -fr test && mkdir test
tbm -w test git license:mit gump:sh
(ls -alh test | grep .git) || exit 1
(ls -alh test | grep .version.sh) || exit 1

rm -fr test && mkdir test
tbm -w test git license:mit gump
(ls -alh test | grep .git) || exit 1
(ls -alh test | grep .version.sh) || exit 1

rm -fr test && mkdir test
tbm -w test git license:mit gump:mh-cbon/emd
(ls -alh test | grep .git) || exit 1
(ls -alh test | grep README.e.md) || exit 1
(cat test/README.e.md | grep "gh-api") || exit 1

rm -fr test && mkdir test
tbm -w test git license:mit emd
(ls -alh test | grep .git) || exit 1
(ls -alh test | grep README.e.md) || exit 1

rm -fr test && mkdir test
tbm -w test git license:mit emd:mh-cbon/emd
(ls -alh test | grep .git) || exit 1
(ls -alh test | grep README.e.md) || exit 1
(cat test/README.e.md | grep "History") || exit 1

rm -fr test && mkdir test
tbm -w test license # ? what happens
(ls -alh test | grep LICENSE) || exit 1
