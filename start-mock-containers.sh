#!/bin/bash
docker network create foobar || true
docker run --rm -d -P -v /var/lib/mysql -e MYSQL_ALLOW_EMPTY_PASSWORD=true mysql:5.7
NGINX_ID=$(docker run --rm -d -P -v /tmp/foobar:/var/www/html nginx:alpine)
docker run --rm -d -P httpd:alpine
TRAEFIK_ID=$(docker run --rm -d -P traefik:maroilles-alpine)
docker run --rm -d -P redis:alpine
docker run --rm -d -P memcached:alpine

docker network connect foobar "$NGINX_ID"
docker network connect foobar "$TRAEFIK_ID"
