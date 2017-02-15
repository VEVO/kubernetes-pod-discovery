#!/bin/bash

export PATH=${PATH}:${GOPATH}/bin

if ! type -P github-release >/dev/null 2>&1 ; then
  go get github.com/aktau/github-release
fi

echo "Tag is ${TAG}"

github-release release \
    --user VEVO \
    --repo kubernetes-pod-discovery \
    --tag ${TAG}

github-release upload \
    --user VEVO \
    --repo kubernetes-pod-discovery \
    --tag ${TAG} \
    --name "kubernetes-pod-discovery" \
    --file kubernetes-pod-discovery
