FROM golang:1.13.0-alpine3.10 as builder

RUN mkdir -p /build 
COPY ./go.mod ./go.sum ./*.go /build/
WORKDIR /build/

RUN CGO_ENABLED=0 go build \
  -installsuffix 'static' \
  -o /build/twiligo .

FROM alpine:3.10.2

COPY --from=builder /build/twiligo /twiligo
EXPOSE 2255/tcp

ENTRYPOINT ["./twiligo"]