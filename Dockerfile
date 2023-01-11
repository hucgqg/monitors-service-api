FROM golang:1.18.4
WORKDIR /opt
COPY monitors-service-api /opt/
COPY config.yml /opt/
RUN mkdir upload
EXPOSE 8080
ENTRYPOINT [ "./monitors-service-api" ]

