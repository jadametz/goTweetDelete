# build
FROM golang:1.13 as build

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /src/
COPY . .
RUN go build -o goTweetDelete

# package
FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /src/goTweetDelete .
CMD ["./goTweetDelete"]
