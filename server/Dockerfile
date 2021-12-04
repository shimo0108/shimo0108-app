FROM golang:latest as dev

WORKDIR /go/src/shimo0108-app/server

COPY . .

RUN go get -u github.com/lib/pq
RUN go get -u google.golang.org/grpc \
    && go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

EXPOSE 9999
CMD ["go", "run", "main.go"]


FROM golang:latest as builder

WORKDIR /go/src/shimo0108-app/server

RUN go get -d -v golang.org/x/net/html

COPY . .

RUN go get -u github.com/lib/pq

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -o app main.go

FROM alpine:latest as prod
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/shimo0108-app/server/app /app
EXPOSE 9999
ENTRYPOINT ["/app"]
