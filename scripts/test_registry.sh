#!/usr/bin/env bash

docker build -t mydocker .
docker run --cap-add="SYS_ADMIN" -e DEBUG \
    mydocker run alpine /bin/echo hey

docker run --cap-add="SYS_ADMIN" -e DEBUG \
    mydocker run ubuntu:latest /bin/echo hey
