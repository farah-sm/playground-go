#!/bin/bash

kubectl api-resources --verbs=list -o name | grep -vE "events" | sort | paste -d, -s - | xargs kubectl get $ns --ignore-not-found --show-kind 2>&1  | grep $1
