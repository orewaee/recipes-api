#!/bin/bash
docker compose --env-file ./config/ports.env -f ./deploy/compose.yaml -p recipes_api up -d
