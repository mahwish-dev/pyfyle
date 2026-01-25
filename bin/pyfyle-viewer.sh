#!/usr/bin/env bash

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd pyfyle || exit
if ! [ -f ".venv/bin.activate" ]; then
  uv venv
fi
source ".venv/bin/activate"
uv run "$SCRIPT_DIR/../main.py" "$@"
