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

rm -fr test && mkdir test
tbm -w test git
(ls -alh test | grep .git) || exit 1

rm -fr test && mkdir test
tbm -w test git license:mit
(ls -alh test | grep .git) || exit 1
(ls -alh test | grep LICENSE) || exit 1

rm -fr test && mkdir test
tbm -w test git license # ? what happens
(ls -alh test | grep .git) || exit 1
