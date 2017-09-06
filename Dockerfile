FROM golang 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go get github.com/go-sql-driver/mysql
RUN ls
RUN go build main.go 
CMD ["/app/main"]
