FROM golang:alpine3.12 AS builder
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o /app/track-info ./cmd/main.go

FROM alpine:3.12

COPY --from=builder /app/track-info /usr/bin/track-info
ARG ARG_LASTFM_APIKEY
ARG ARG_LASTFM_URL
ENV LASTFM_APIKEY=$ARG_LASTFM_APIKEY
ENV LASTFM_URL=$ARG_LASTFM_URL

ENTRYPOINT ["track-info"]
