FROM alpine:edge
ADD localhost.crt /etc/ssl/certs/
ADD localhost.key /etc/ssl/certs/
ADD main /
RUN apk add logrotate
RUN mkdir /var/log/go-service
COPY go-service /etc/logrotate.d/go-service
CMD ["/main"]