#!/bin/bash
set -uex
# json=`printf '{"name": "%s", "email": "%s", "password": "%s"}' $1 $2 $3`
# echo $json
curl -v -X POST -H 'Content-type: application/json' -d "`printf '{"name": "%s", "email": "%s", "password": "%s"}' $1 $2 $3`" localhost:8888/api/users
