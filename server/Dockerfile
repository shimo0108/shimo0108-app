FROM golang:latest as dev

WORKDIR /go/src/shimo0108-app/server
RUN mkdir /tmp/ssm
RUN wget https://s3.amazonaws.com/ec2-downloads-windows/SSMAgent/latest/debian_amd64/amazon-ssm-agent.deb -O /tmp/ssm/amazon-ssm-agent.deb \
    && sudo dpkg -i /tmp/ssm/amazon-ssm-agent.deb \
    && cp /etc/amazon/ssm/seelog.xml.template /etc/amazon/ssm/seelog.xml

COPY . .

RUN go get -u github.com/lib/pq
RUN go get -u google.golang.org/grpc \
    && go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
ENTRYPOINT [ "/docker-entrypoint.sh" ]

EXPOSE 9999
CMD ["go", "run", "main.go"]
