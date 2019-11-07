FROM golang

RUN go get github.com/astaxie/beego && go get github.com/beego/bee &&go get github.com/go-sql-driver/mysql

WORKDIR src/flight

COPY * .

CMD bee run