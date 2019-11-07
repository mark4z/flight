FROM golang

RUN go get -d -v ./...

CMD bee run