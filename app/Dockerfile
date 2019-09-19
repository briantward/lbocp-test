FROM registry.redhat.io/rhel8/go-toolset

COPY main.go ./

USER root

RUN go build main.go

EXPOSE 8080

CMD ["./main"] 
