#!bin/bash

go build -ldflags="-s -w" -o bin/toys main.go

bindir=""

cp -r bin $bindir