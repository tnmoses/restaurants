FROM golang:latest
RUN mkdir /app
ADD ./main.go /app/
WORKDIR /app
EXPOSE 8080
RUN go build -o restaurants .
CMD ["/app/restaurants"]
