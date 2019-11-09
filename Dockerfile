FROM golang

RUN go get github.com/astaxie/beego && go get github.com/beego/bee &&go get github.com/go-sql-driver/mysql && go get github.com/astaxie/beego/cache && && go get github.com/astaxie/beego/cache/redis && go get github.com/sideshow/apns2

WORKDIR src/flight

COPY / .

CMD bee run