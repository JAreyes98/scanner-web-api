FROM ubuntu:latest
LABEL authors="jdreyes"

ENTRYPOINT ["top", "-b"]