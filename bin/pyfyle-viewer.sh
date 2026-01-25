#!/usr/bin/env bash

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
pwd
cd "$SCRIPT_DIR/pyfyle" || exit
uv venv
source "$SCRIPT_DIR/pyfyle/.venv/bin/activate"
cd ..

uv run "$SCRIPT_DIR/../main.py" "$@"
