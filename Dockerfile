FROM golang:1.6
WORKDIR /go/src/github.com/EtienneR/go_crypto_generator
ADD . /go/src/github.com/EtienneR/go_crypto_generator
RUN go install github.com/EtienneR/go_crypto_generator
ENTRYPOINT /go/bin/go_crypto_generator
EXPOSE 3000