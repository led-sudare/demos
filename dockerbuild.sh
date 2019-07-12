#!/bin/sh
cname=`cat ./cname`
docker build ./ -t $cname
docker run -t --init --name $cname -v `pwd`:/work/ -p 5003:5003 $cname