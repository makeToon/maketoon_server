#FROM golang:alpine
#
#WORKDIR /go/src/maketoon
#COPY . .
#
#RUN apk update && \
#  apk add git && \
#  go get github.com/cespare/reflex
#
#EXPOSE 9999
#CMD ["reflex", "-c", "reflex.conf"]

FROM golang:latest AS builder
WORKDIR /go/src/maketoon
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w" -a -o /main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /main ./
RUN chmod +x ./main
EXPOSE 8080
CMD ./main