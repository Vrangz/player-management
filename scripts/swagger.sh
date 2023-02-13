#!/bin/bash

swagger -q generate spec -m -w ./player-manager/internal -o ./swagger-ui-dist/swagger.yaml 