#!/usr/bin/env bash

docker build -t mydocker .
docker run --cap-add="SYS_ADMIN" -e DEBUG \
    mydocker pull ubuntu:latest
