FROM golang:1.18.4
WORKDIR /opt
RUN go build -o monitors-service-api main.go
COPY monitors-service-api /opt/
COPY config.yml /opt/
RUN mkdir upload
EXPOSE 8080
ENTRYPOINT [ "./monitors-service-api" ]
