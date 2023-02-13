#!/bin/bash

NAME="xo"

docker image build -t ${NAME} -f ./deployment/xo.Dockerfile .

docker compose \
    -f ./deployment/db-compose.yaml \
    -f ./deployment/xo-compose.yaml \
    up -d
