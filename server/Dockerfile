FROM golang:latest as dev

WORKDIR /go/src/shimo0108-app/server

COPY . .

RUN go get -u github.com/lib/pq
RUN go get -u google.golang.org/grpc \
    && go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

EXPOSE 9999
CMD ["go", "run", "main.go"]
