@echo off
docker compose -f ./deploy/compose.yaml -p recipes_api up -d
