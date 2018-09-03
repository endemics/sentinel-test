FROM golang:1.10.2-alpine3.7 as builder
RUN apk add --no-cache curl && \
    curl -L https://releases.hashicorp.com/sentinel/0.3.1/sentinel_0.3.1_linux_amd64.zip --output sentinel.zip && \
    unzip sentinel.zip && \
    mv sentinel /usr/local/bin/sentinel && \
    rm sentinel.zip
COPY . /go/src/sentinel-test
RUN go install sentinel-test

FROM alpine:3.7
COPY --from=builder /usr/local/bin/sentinel /usr/local/bin/sentinel
COPY --from=builder /go/bin/sentinel-test /usr/local/bin/sentinel-test
COPY . /src
WORKDIR /src
ENTRYPOINT ["/usr/local/bin/sentinel","apply","-config=config.json","sentinel.policy"]
