FROM golang:latest as dev

WORKDIR /go/src/shimo0108-app/server

COPY . .

RUN go get -u github.com/lib/pq
RUN go get -u google.golang.org/grpc \
    && go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

EXPOSE 9999
CMD ["go", "run", "main.go"]


FROM golang:latest as prod
RUN apt-get update && apt-get install -y sudo

RUN mkdir /tmp/ssm
RUN wget https://s3.amazonaws.com/ec2-downloads-windows/SSMAgent/latest/debian_amd64/amazon-ssm-agent.deb -O /tmp/ssm/amazon-ssm-agent.deb \
    && sudo dpkg -i /tmp/ssm/amazon-ssm-agent.deb \
    && cp /etc/amazon/ssm/seelog.xml.template /etc/amazon/ssm/seelog.xml


WORKDIR /go/src/shimo0108-app/server

COPY . .

RUN go get -u github.com/lib/pq
RUN go get -u google.golang.org/grpc \
    && go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

COPY docker-entrypoint.sh /
ENTRYPOINT [ "/docker-entrypoint.sh" ]
