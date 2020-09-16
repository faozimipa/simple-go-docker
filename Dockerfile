FROM golang:1.12.0-alpine3.9
RUN mkdir /appgo
ADD . /appgo
WORKDIR /appgo
RUN go build -o main .
CMD ["/appgo/main"]