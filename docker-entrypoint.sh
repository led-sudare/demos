#!/bin/sh

#newdemo
echo "*** docker-entorypoint.sh Start... ***"
cname=`cat ./cname`
cd moridemo
make clean
make
#chmod a+x moridemo

cd ../sudare_contents
go build

cd ..
go build
exec ./demos

#exec ./moridemo/demo
#exec ./sudare_contents/sudare_contents