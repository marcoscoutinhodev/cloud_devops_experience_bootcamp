#!/bin/bash

docker build -t marcoscoutinhodev/backend:1.0 ./backend/ && \
  docker push marcoscoutinhodev/backend:1.0


docker build -t marcoscoutinhodev/backend-db:1.0 ./db/ && \
  docker push marcoscoutinhodev/backend-db:1.0


kubectl apply -f ./backend/k8s/ && \
kubectl apply -f ./db/k8s/
