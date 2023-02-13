#!/bin/bash

docker compose \
    -f ./deployment/food-consumer-compose.yaml \
    -f ./deployment/player-manager-compose.yaml \
    -f ./deployment/xo-compose.yaml \
    -f ./deployment/db-compose.yaml \
    down

docker volume rm deployment_db
