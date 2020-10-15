#!/usr/bin/env bash

kubectl create -f db-service.yaml,db-deployment.yaml,bulletin-api-service.yaml,bulletin-api-deployment.yaml