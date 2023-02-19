#!/usr/bin/env bash

docker build -t mydocker . && docker run --cap-add="SYS_ADMIN" mydocker run ubuntu:latest /usr/local/bin/docker-explorer ls /some_dir
