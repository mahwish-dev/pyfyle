#!/usr/bin/env bash

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd pyfyle || exit
uv venv
source ".venv/bin/activate"
uv run "$SCRIPT_DIR/../main.py" "$@"
