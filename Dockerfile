FROM golang

ADD . /usr/src/ExpensiveProduct

WORKDIR /usr/src/ExpensiveProduct

# Build the Go app
RUN go build -o main .

# Command to run the executable
CMD ["./main"]