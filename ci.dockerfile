ARG GO_VERSION
FROM golang:${GO_VERSION}
WORKDIR /go/src/github.com/geowa4/base-go
COPY . .
RUN make deps
RUN make
RUN ls
