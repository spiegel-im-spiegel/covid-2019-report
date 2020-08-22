#!/bin/bash
./build-chart.sh || exit 1
comment="Update report in $(date -u '+%F')"
git add --all || exit 1
git commit -v -m "$comment" || exit 1
git push -u origin master || exit 1
