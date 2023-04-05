#!/usr/bin/env bash

set -euo pipefail
shopt -s nullglob

if [ "$#" -ne 1 ]; then
    echo "usage: ./migrate.sh DATABASE_URL"
    exit 1
fi

for fname in db/*.up.sql; do
    psql $1 < $fname
done
