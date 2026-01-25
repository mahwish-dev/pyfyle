#!/usr/bin/env bash

cd pyfyle/site || exit
hugo server --disableFastRender -D -t re-terminal
cd ..
cd ..
