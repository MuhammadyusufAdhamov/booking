# Multistage dockerfile (it is for minimizing docker image size)

# -- Build stage --
FROM golang:1.19.1-alpine3.16 AS builder

# workdir is the current working directory inside docker image
# all dockerfile instructions will be executed inside workdir
WORKDIR /app

# first dot means that copy everything from current folder (booking folder)
# second dot is the current working directory inside the image (/app folder)
COPY . .

RUN apk add curl
RUN go build -o main cmd/main.go
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

# -- Run stage --
FROM alpine:3.16

WORKDIR /app
RUN mkdir media

# copying main binary file to workdir
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate
COPY migrations ./migrations
COPY templates ./templates

EXPOSE 8000

CMD ["/app/main"]


# create docker network
# docker network create booking-network
# docker network connect booking-network postgres
# docker network connect booking-network redis-cli
# docker run --env-file ./.env --name bookingApp --network booking-network -p 8000:8000 -d booking:latest
