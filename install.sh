#!/bin/sh

BINARY='/usr/local/bin'

echo "Building index_drive"
go build index_drive.go

echo "Installing index_drive to $BINARY"
install -v index_drive $BINARY

echo "Removing the build"
rm index_drive
