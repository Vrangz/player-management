#!/bin/bash

PM_NAME="player-manager"
FC_NAME="food-consumer"

docker image build -t ${PM_NAME} -f ./deployment/player-manager.Dockerfile .
docker image build -t ${FC_NAME} -f ./deployment/food-consumer.Dockerfile .

docker compose \
    -f ./deployment/db-compose.yaml \
    -f ./deployment/player-manager-compose.yaml \
    -f ./deployment/food-consumer-compose.yaml \
    up -d
