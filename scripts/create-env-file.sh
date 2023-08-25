#!/usr/bin/env bash

DEPLOY_DIR=$(pwd)/deploy
CONFIG_DIR=$DEPLOY_DIR/config/

envsubst < $CONFIG_DIR/env.tfvars.template > $DEPLOY_DIR/env.auto.tfvars

cat $DEPLOY_DIR/env.auto.tfvars
