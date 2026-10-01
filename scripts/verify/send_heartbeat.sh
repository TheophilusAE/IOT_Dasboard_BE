#!/usr/bin/env bash
set -euo pipefail

HOST="${HOST:-http://localhost:8080}"
DEVICE_UID="${DEVICE_UID:-TEST-ESP32-0001}"

curl -sS -X POST "$HOST/api/v1/ingest/heartbeat" \
  -H "X-Device-UID: $DEVICE_UID" \
  -H "Content-Type: application/json" \
  -d '{
    "uptime_seconds": 120,
    "rssi": -55,
    "free_heap_bytes": 145200,
    "firmware_version": "test-1.0"
  }'
echo
