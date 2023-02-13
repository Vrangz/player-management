#!/bin/bash

xo --verbose schema postgres://$DB_USER:$DB_PASSWORD@$DB_HOST/$DB_NAME?sslmode=disable -o /xo
