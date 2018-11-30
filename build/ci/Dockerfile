ARG GO_VERSION
FROM golang:${GO_VERSION}
RUN mkdir -p /ci
WORKDIR /ci
COPY . .
RUN make deps
RUN make
