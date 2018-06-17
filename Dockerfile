FROM golang:latest
COPY ./app /go/src/github.com/tnmoses/restaurants
WORKDIR /go/src/github.com/tnmoses/restaurants
EXPOSE 8080
RUN go get ./
RUN go build main.go handlers.go utils.go restaurant.go -o restaurants .
CMD ["/go/src/github.com/tnmoses/restaurants/restaurants"]
