#!/bin/bash

curl -v -X POST -H 'Content-type: application/json' -d '{"name": "hoge", "email": "aaa@example.com", "password": "aaaa1111"}' localhost:8888/api/users
