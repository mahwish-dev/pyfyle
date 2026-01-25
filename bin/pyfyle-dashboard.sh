#!/usr/bin/env bash

cd pyfyle/site || exit
hugo server -D -t re-terminal
cd ..
cd ..
