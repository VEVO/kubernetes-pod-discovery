#!/bin/bash

set -ex

go get github.com/Masterminds/glide && \
  glide install --strip-vendor
