FROM alpine
CMD [ "/service" ]
RUN apk --no-cache add ca-certificates
ARG SERVICE_NAME
ENV SERVICE_NAME=${SERVICE_NAME}
COPY ./${SERVICE_NAME}-linux-amd64 /service
COPY ./${SERVICE_NAME}-linux-amd64.md5 /service.md5
COPY ./${SERVICE_NAME}-linux-amd64.sha256 /service.sha256
