FROM golang:1.8

WORKDIR /
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]