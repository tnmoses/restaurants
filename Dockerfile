FROM golang:latest
COPY ./app /go/src/github.com/tnmoses/restaurants/app
WORKDIR /go/src/github.com/tnmoses/restaurants/app
EXPOSE 8080
RUN go get ./
RUN go build -o restaurants .
CMD ["/go/src/github.com/tnmoses/restaurants/app/restaurants"]
