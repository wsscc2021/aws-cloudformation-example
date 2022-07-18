#!/bin/bash
#WORKDIR
cd /opt/app
if [ ! -d .venv ]; then
    pypy3 -m venv .venv
fi
pypy3 -m venv .venv --clear
.venv/bin/pip3 install -r requirements.txt