#!/bin/sh

ENVIRONMENT=${1:-dev}

exchange_rate_env_path="../envs/$ENVIRONMENT/.env.exchange-rate"

if [ "$ENVIRONMENT" = "prod" ]; then
	echo "Starting production server"
	echo "Path to exchange rate env file: $exchange_rate_env_path"
  	export APP_ENV_FILE=.env.app.prod
else
	echo "Starting development server"
	echo "Path to exchange rate env file: $exchange_rate_env_path"
fi
