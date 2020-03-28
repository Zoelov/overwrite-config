# Base image: https://hub.docker.com/_/golang/
FROM ubuntu:14.04
MAINTAINER wanghao <wanghao000.cool@163.com>

COPY merge-config /usr/local/bin/
