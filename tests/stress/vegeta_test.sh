#!/usr/bin/env bash

set -euo pipefail
echo "GET http://localhost:8080/ping" | vegeta attack -duration=15s -rate=200 | vegeta report
# echo "GET http://localhost:8080/users" | vegeta attack -duration=5s -rate=1000 | vegeta report
