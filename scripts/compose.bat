@echo off
docker compose -f ./deploy/compose.yaml -p recipes-api up -d
