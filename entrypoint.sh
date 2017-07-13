#!/bin/bash
set -e

case "$1" in
    develop)
        echo "Running Develop"
        echo -e "$GCLOUD_CREDENTIALS" | base64 -d > credentials.json
        exec gin -p $PORT run main.go
        ;;
    start)
        echo "Running Start"
        echo "$GCLOUD_CREDENTIALS" | base64 -d > credentials.json
        exec go run main.go
        ;;
    startkube)
        echo "Running Start in Kube Cluster"
        echo -e "$GCLOUD_CREDENTIALS" > credentials.json
        exec go run main.go
        ;;
    *)
        exec "$@"
esac
