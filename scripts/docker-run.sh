#!/usr/bin/env bash

docker run \
  -it \
  --rm \
  -p 8080:8080 \
	--env SYNOPTIC_API_TOKEN=$(SYNOPTIC_API_TOKEN) \
  colevoss/temp-blanket-backend
