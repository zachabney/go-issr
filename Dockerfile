FROM golang
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go build main.go
CMD ["/app/main"]

EXPOSE 8080