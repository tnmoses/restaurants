FROM golang:latest
COPY ./ /go/src/github.com/tnmoses/restaurants
WORKDIR /go/src/github.com/tnmoses/restaurants
EXPOSE 8080
RUN go get ./
RUN go build -o restaurants main.go handlers.go utils.go restaurant.go
CMD ["/go/src/github.com/tnmoses/restaurants/restaurants"]
