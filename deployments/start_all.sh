#!/bin/bash

kubectl -n go-jwt-mysql-api apply -f mysql-secret.yaml
kubectl -n go-jwt-mysql-api apply -f mysql-deployment.yaml
kubectl -n go-jwt-mysql-api apply -f app-mysql-deployment.yaml
