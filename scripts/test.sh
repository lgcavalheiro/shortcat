#!/bin/sh

while IFS= read -r folder; do
    lib="github.com/lgcavalheiro/shortcat/$folder"
    go test "$lib" -v -coverprofile=/tmp/vscode-go5OhECo/go-code-cover "$lib"
done <<-EOF
model
server
util
	EOF
