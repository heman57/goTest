FROM golang
RUN go get "github.com/go-sql-driver/mysql"

RUN mkdir /app 
ADD . /app/
WORKDIR /app 

EXPOSE 80

CMD ["/app/main"]



