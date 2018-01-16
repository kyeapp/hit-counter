#!/bin/bash
go fmt
GOPATH="C:\Users\KevinD\Documents\GitHub\server" go build -v -x
if [ $? -eq 0 ]
then
	echo "build successful"
else 
	echo "build failed."
fi
