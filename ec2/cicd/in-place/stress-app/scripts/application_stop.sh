#!/bin/bash
if [ -e /opt/app/pid ]; then
    kill -15 $(cat /opt/app/pid)
fi
