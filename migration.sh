#!/bin/bash
source config.env

sleep 20 && goose -dir "${MIGRATION_DIR}" postgres "${MIGRATION_DSN}" up -v