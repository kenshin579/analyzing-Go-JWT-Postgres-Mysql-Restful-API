#!/bin/bash

kubectl -n go-jwt-mysql-api delete -f app-mysql-deployment.yaml
kubectl -n go-jwt-mysql-api delete -f mysql-deployment.yaml
kubectl -n go-jwt-mysql-api delete -f mysql-secret.yaml