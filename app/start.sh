#!/bin/bash
trap "rm server;kill 0" EXIT
go build -o server
./server -port=8081 &
./server -port=8082 &
./server -port=8083 -api=true &

sleep 2
echo ">>> start test"
curl "http://localhost:9999/api?key=Sam" &
curl "http://localhost:9999/api?key=Sam" &
curl "http://localhost:9999/api?key=Sam" &
wait