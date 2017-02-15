#!/bin/bash

set -ex

declare -r binary_name="${BINARY_NAME:-kubernetes-pod-discovery}"

go build -v -o ${binary_name} .
