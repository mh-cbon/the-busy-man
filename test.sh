#!/bin/sh

set -e
# set -x

OLDPWD=`pwd`
go install

rm -fr ~/test-tbm && mkdir -p ~/test-tbm
cd ~/test-tbm
TESTDIR=`pwd`
cd $OLDPWD

rm -fr $TESTDIR && mkdir -p $TESTDIR

alias tbm='the-busy-man'

tbm -h
tbm -help
tbm -v
tbm -version
tbm -l
tbm -h git

export VERBOSE=y
export VERBOSE=""

rm -fr $TESTDIR && mkdir -p $TESTDIR
tbm -w $TESTDIR git
echo ""
(ls -alh $TESTDIR | grep .git) || exit 1

echo ""
echo ""

rm -fr $TESTDIR && mkdir -p $TESTDIR
tbm -w $TESTDIR git license:mit
echo "==>"
(ls -alh $TESTDIR | grep .git) || exit 1
(ls -alh $TESTDIR | grep LICENSE) || exit 1

echo ""
echo ""

rm -fr $TESTDIR && mkdir -p $TESTDIR
tbm -w $TESTDIR git license:mit git:init+commit changelog
echo "==>"
(ls -alh $TESTDIR | grep .git) || exit 1
(ls -alh $TESTDIR | grep LICENSE) || exit 1
(ls -alh $TESTDIR | grep change.log) || exit 1

echo ""
echo ""

rm -fr $TESTDIR && mkdir -p $TESTDIR
tbm -w $TESTDIR git license:mit golang
echo "==>"
(ls -alh $TESTDIR | grep .git) || exit 1
(ls -alh $TESTDIR | grep main.go) || exit 1

echo ""
echo ""

rm -fr $TESTDIR && mkdir -p $TESTDIR
tbm -w $TESTDIR git license:mit gump:sh
echo "==>"
(ls -alh $TESTDIR | grep .git) || exit 1
(ls -alh $TESTDIR | grep .version.sh) || exit 1

echo ""
echo ""

rm -fr $TESTDIR && mkdir -p $TESTDIR
tbm -w $TESTDIR git license:mit gump
echo "==>"
(ls -alh $TESTDIR | grep .git) || exit 1
(ls -alh $TESTDIR | grep .version.sh) || exit 1

echo ""
echo ""

rm -fr $TESTDIR && mkdir -p $TESTDIR
tbm -w $TESTDIR git license:mit gump:mh-cbon/emd
echo "==>"
(ls -alh $TESTDIR | grep .git) || exit 1
(ls -alh $TESTDIR | grep .version.sh) || exit 1
(cat $TESTDIR/.version.sh | grep "gh-api") || exit 1

echo ""
echo ""

rm -fr $TESTDIR && mkdir -p $TESTDIR
tbm -w $TESTDIR git license:mit emd
echo "==>"
(ls -alh $TESTDIR | grep .git) || exit 1
(ls -alh $TESTDIR | grep README.e.md) || exit 1

echo ""
echo ""

rm -fr $TESTDIR && mkdir -p $TESTDIR
tbm -w $TESTDIR git license:mit emd:mh-cbon/emd
echo "==>"
(ls -alh $TESTDIR | grep .git) || exit 1
(ls -alh $TESTDIR | grep README.e.md) || exit 1
(cat $TESTDIR/README.e.md | grep "History") || exit 1

echo ""
echo ""

rm -fr $TESTDIR && mkdir -p $TESTDIR
tbm -w $TESTDIR git go glide
echo "==>"
(ls -alh $TESTDIR | grep .git) || exit 1
(ls -alh $TESTDIR | grep main.go) || exit 1
(ls -alh $TESTDIR | grep glide) || exit 1

echo ""
echo ""

rm -fr $TESTDIR && mkdir -p $TESTDIR
tbm -w $TESTDIR license | grep "missing license"
# echo "==>"

echo ""
echo ""

echo "All Good!"
