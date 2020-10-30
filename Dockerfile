FROM golang:1.13-alpine as build
WORKDIR $GOPATH/src/ERT/
ADD . $GOPATH/src/ERT/
RUN go mod vendor && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main $GOPATH/src/ERT/cmd/webhook-server



#FROM gcr.io/distroless/base
FROM alpine:3.9
LABEL maintainer="Brian Liu <jay@sparkflow.top>"
VOLUME ['/log']
COPY --from=build /go/src/ERT/main /main
ENTRYPOINT ["/main"]