#!/usr/bin/env bash

set -ueo pipefail

function abort {
  echo "$@" 1>&2
  exit 1
}

cmd="$1"
case "$cmd" in
  up)
    migrate -database 'mysql://root:secret@tcp(127.0.0.1:3306)/imagechat' -path mysql/migrations/ up
  ;;
  down)
    migrate -database 'mysql://root:secret@tcp(127.0.0.1:3306)/imagechat' -path mysql/migrations/ down
  ;;
  version)
    migrate -database 'mysql://root:secret@tcp(127.0.0.1:3306)/imagechat' -path mysql/migrations/ version
  ;;
  create)
    migration_name="$2"
    migrate create -ext sql -dir mysql/migrations -seq "$migration_name"
  ;;
  *)
    abort "no such command: " "$cmd"
  ;;
esac
