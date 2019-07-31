#!/bin/sh

echo "*** docker-entorypoint.sh Start... ***"
cname=`cat ./cname`
cd moridemo
rm ./moridemo
make

cd ../sudare_contents
rm ./sudare_contents
go build

cd ..
rm ./demos
go build
exec ./demos
