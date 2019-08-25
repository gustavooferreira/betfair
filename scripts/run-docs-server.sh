#!/bin/bash


function _godoc() {
    [ ! -f "$(pwd)/go.mod" ] && echo "error: go.mod not found" || module=$(awk 'NR==1{print $2}' go.mod) && docker run --rm -e "GOPATH=/tmp/go" -p 6060:6060 -v $(pwd):/tmp/go/src/$module golang:1.12.6 /bin/bash -c "awk 'END{print \"http://\"\$1\":6060/pkg/$module\"}' /etc/hosts && godoc -http=:6060";
};

_godoc
