# Stage 1. Build the binary
FROM golang:1.11

# add a non-privileged user
RUN useradd -u 10001 myapp

RUN mkdir -p /go/src/github.com/middlehut/workshop.k8s
ADD . /go/src/github.com/middlehut/workshop.k8s
WORKDIR /go/src/github.com/middlehut/workshop.k8s

# build the binary with go build
RUN CGO_ENABLED=0 go build \
	-o bin/workshop.k8s github.com/middlehut/workshop.k8s/cmd/workshop.k8s

# Stage 2. Run the binary
FROM scratch

ENV PORT 8080
ENV DIAG_PORT 8585

COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=0 /etc/passwd /etc/passwd
USER myapp

COPY --from=0 /go/src/github.com/middlehut/workshop.k8s/bin/workshop.k8s /workshop.k8s
EXPOSE $PORT
EXPOSE $DIAG_PORT

CMD ["/workshop.k8s"]