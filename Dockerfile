##  Building Stage
FROM golang:1.13.0-alpine as builder

RUN apk add --no-cache make git \
  && go get github.com/oxequa/realize


WORKDIR /go/src/github.com/hukurou-s/user-auth-api-with-jwt
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
COPY . .
RUN make build


## Running Stage

FROM alpine:3.10.2
WORKDIR /app
COPY --from=builder /go/src/github.com/hukurou-s/user-auth-api-with-jwt/user-auth-api-with-jwt /app/user-auth-api
EXPOSE 1322

CMD ["/app/user-auth-api"]
