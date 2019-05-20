# builder image
FROM golang:1.12-alpine3.9 as builder
MAINTAINER Bouwe Ceunen <bouweceunen@gmail.com>

RUN apk -U add git
ENV CGO_ENABLED 0
COPY webhook /webhook
RUN GOPATH=/webhook
RUN cd /webhook && go build -o /bin/webhook

# final image
FROM lachlanevenson/k8s-kubectl
MAINTAINER Bouwe Ceunen <bouweceunen@gmail.com>
WORKDIR /hook
COPY --from=builder /bin/webhook /hook/webhook
COPY --from=builder /webhook/argo.yml /hook/argo.yml
ENTRYPOINT [ "" ]
CMD ["/hook/webhook"]