#!bin/bash

repodir=""

go build -ldflags="-s -w" -o $repodir/bin/toys $repodir/main.go

bindir=""

cp -r bin $bindir