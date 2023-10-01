# First stage, building application
FROM golang:alpine AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app/

RUN go mod download
RUN go build cmd/main/main.go

# Last stage : Creating final container
FROM alpine
WORKDIR /
COPY --from=builder /app/main /main
ENTRYPOINT /main
