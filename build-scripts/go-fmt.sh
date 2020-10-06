#!/bin/bash

# Adapted from https://github.com/Jerome1337/gofmt-action

cd $GITHUB_WORKSPACE

GOFMT_OUTPUT="$(gofmt -l . 2>&1)"

if [ -n "$GOFMT_OUTPUT" ]; then
  echo "All the following files are not correctly formatted"
  echo "${GOFMT_OUTPUT}"

  exit 1
fi

echo "::set-output name=gofmt-output::Gofmt step succeed"