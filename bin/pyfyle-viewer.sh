#!/usr/bin/env bash

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
pwd
cd pyfyle || exit
pwd
uv venv
source ".venv/bin/activate"
echo "$@"
uv run "$SCRIPT_DIR/../main.py" "$@"
