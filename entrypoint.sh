#!/bin/bash
set -e

case "$1" in
    develop)
        echo "Running Develop"
        exec gin -p $PORT run main.go
        ;;
    start)
        echo "Running Start"
        exec go run main.go
        ;;
    *)
        exec "$@"
esac
