FROM golang:1.9.3-alpine

LABEL maintainer='Jumia TechOps <techops@jumia.com>'

ENV \
    DEPS='' \
    PKGS='' \
    \
    STATUSCAKE_ENDPOINT='' \
    STATUSCAKE_AUTH='SECRET'

WORKDIR /app

COPY app/ /app
