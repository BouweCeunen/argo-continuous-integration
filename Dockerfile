# builder image
FROM golang:1.12-alpine3.9 as builder
MAINTAINER Bouwe Ceunen <bouweceunen@gmail.com>

ENV CGO_ENABLED 0
RUN apk --no-cache add git
RUN go get github.com/golang/dep/cmd/dep
WORKDIR /go/src/github.com/bouweceunen/webhook
COPY . .
RUN dep ensure -vendor-only
RUN go test -v ./...
ENV GOARCH amd64
RUN go build -o /bin/webhook -v -ldflags "-X main.version=$(git describe --tags --always --dirty) -w -s"
RUN mkdir /tmp/result/ && cp /bin/webhook /tmp/result/webhook

# final image
FROM gcr.io/distroless/base
MAINTAINER Bouwe Ceunen <bouweceunen@gmail.com>
COPY --from=builder /tmp/result/ /
CMD ["/webhook"]