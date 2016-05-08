#!/bin/bash
while true; do
  find -name *.php | shuf | head -n 1 | xargs cat | shuf | tail -n 1
  sleep 1
done
