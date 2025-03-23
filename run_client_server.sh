#!/bin/bash
echo "####START####"

echo "Running server..."
nohup go run server/server.go &

echo "Running client..."
go run client/client.go
killall server
echo "####END####"


