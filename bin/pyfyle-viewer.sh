#!/usr/bin/env bash

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd ..
uv venv
source ".venv/bin/activate"
cd bin/ || exit

uv run "$SCRIPT_DIR/../main.py" "$@"
