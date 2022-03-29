FROM golang

ADD . /usr/src/ExpensiveProduct

WORKDIR /usr/src/ExpensiveProduct

RUN go get github.com/codegangsta/gin

RUN go build

CMD bash -c "cd /tmp && gin --path /usr/src/ExpensiveProduct"
