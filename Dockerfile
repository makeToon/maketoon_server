#FROM golang:alpine
#
#WORKDIR /go/src/maketoon
#COPY . .
#
#RUN apk add --no-cache git mercurial
#RUN go get -d -v ./...
#RUN go install -v ./...
#RUN apk del git mercurial
#
#ENTRYPOINT ["go", "run", "main.go"]

FROM golang:alpine

WORKDIR /go/src/maketoon
COPY . .

RUN apk update && \
  apk add git && \
  go get github.com/cespare/reflex

EXPOSE 9999
CMD ["reflex", "-c", "reflex.conf"]