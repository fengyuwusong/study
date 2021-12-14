#!/bin/bash
trap "rm server;kill 0" EXIT

# 跳转到文件所在路径
CURPATH=$(cd "$(dirname "$0")"; pwd)
cd $CURPATH

go build -o server
./server -port=8001 &
./server -port=8002 &
./server -port=8003 -api=1 &

sleep 2
echo ">>> start test"
curl "http://localhost:9999/api?key=Tom" &
curl "http://localhost:9999/api?key=Tom" &
curl "http://localhost:9999/api?key=Tom" &

wait