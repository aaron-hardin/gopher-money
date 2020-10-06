#!/bin/bash

# Adapted from https://github.com/Jerome1337/gofmt-action

cd $GITHUB_WORKSPACE

if [ "$(gofmt -l . | wc -l)" -gt 0 ]; then
  echo "Issues with formatting:"
  gofmt -e .
  exit 1
fi

echo "::set-output name=gofmt-output::Gofmt step succeed"