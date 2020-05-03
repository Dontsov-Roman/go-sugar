#!/bin/sh

set -e
host="$1"
shift

until nc -z -v -w30 $host 3306
do
  echo "Waiting for database connection..."
  echo $host
  # wait for 5 seconds before check again
  sleep 5
done