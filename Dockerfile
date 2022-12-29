FROM golang:latest AS builder
WORKDIR /build
ADD . .

