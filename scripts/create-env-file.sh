#!/usr/bin/env bash

CONFIG_DIR=$(pwd)/deploy/config/

envsubst < $CONFIG_DIR/env.tfvars.template > $CONFIG_DIR/env.auto.tfvars
