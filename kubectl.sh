#!/bin/env bash

kubectl apply -f db-service.yaml,db-deployment.yaml,bulletinapi-service.yaml,bulletinapi-deployment.yaml