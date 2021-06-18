FROM golang:alpine3.13

WORKDIR /go/src/URLShortener

COPY ./ ./

RUN go get -d -v ./...
RUN go install -v ./...

CMD [ "/go/bin/server", "-config-path", "./config/config.docker.yml" ]