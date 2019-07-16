#!/bin/sh

echo "*** docker-entorypoint.sh Start... ***"
cname=`cat ./cname`
cd moridemo
make clean
make

cd ../sudare_contents
rm ./sudare_contents
go build

cd ..
rm ./demos
go build
exec ./demos
