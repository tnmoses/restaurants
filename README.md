# Restaurants
A REST API built in Go backed by BoltDB, a pure Go database.
## Table of Contents

- [Getting Started](#getting-started)
- [API Documentation](#api-documentation)
  - [Create new restaurant](#create-new-restaurant)
  - [Get all restaurants](#get-all-restaurants)
  - [Get one restaurant](#get-one-restaurant)
  - [Update restaurant](#update-restaurant)
  - [Delete restaurant](#delete-restaurant)
  - [Health Check](#health-check)

## Getting Started
```bash
git clone github.com/tnmoses/restaurants.git
cd restaurants
docker build -t restaurants -f Dockerfile .
docker run -p 8080:8080 restaurants
```
## API Documentation
### Create new restaurant
POST /restaurants
```JSON
{
    "Name": "ACME Dining",
    "Phone": "+491773456543",
    "Cuisines": "mediterranean, mild",
    "Address": "Downtown, Tasteville",
    "Description": "Excellent food"
}
```

##### Success response
HTTP Status Code: 200

```JSON
{
    "ID": 1,
    "Name": "ACME Dining",
    "Phone": "+491773456543",
    "Cuisines": "mediterranean, mild",
    "Address": "Downtown, Tasteville",
    "Description": "Excellent food"
}
```
##### Failure response
HTTP Status Code: 400
```JSON
{
    "error": "Invalid request payload"
}
```
### Get all restaurants
GET /restaurants

##### Success response
HTTP Status Code: 200

```JSON
[
  {
    "ID": 1,
    "Name": "ACME Dining",
    "Phone": "+491773456543",
    "Cuisines": "mediterranean, mild",
    "Address": "Downtown, Tasteville",
    "Description": "Excellent food"
  }
]
```
### Get one restaurant
GET /restaurants/:id

##### Success response
HTTP Status Code: 200

```JSON
{
  "ID": 1,
  "Name": "ACME Dining",
  "Phone": "+491773456543",
  "Cuisines": "mediterranean, mild",
  "Address": "Downtown, Tasteville",
  "Description": "Excellent food"
}
```
##### Failure Response
HTTP Status Code: 404
```JSON
{
  "error": "Restaurant not found"
}
```
### Update restaurant
PUT /restaurants/:id
```JSON
{
  "Description": "Good service"
}
```

##### Success response
HTTP Status Code: 200

```JSON
{
    "ID": 1,
    "Name": "ACME Dining",
    "Phone": "+491773456543",
    "Cuisines": "mediterranean, mild",
    "Address": "Downtown, Tasteville",
    "Description": "Good service"
}
```
Only the fields sent in the request payload will be updated; all other fields remain the same. Sending a field with an empty string is allowed. Updating the primary key is not allowed.
##### Failure Response
HTTP Status Code: 400
```JSON
{
  "error": "Invalid request payload"
}
```
### Delete restaurant
DELETE /restaurants/:id

##### Success response
HTTP Status Code: 204 (No content)

##### Failure Response
HTTP Status Code: 404
```JSON
{
  "error": "Restaurant not found"
}
```
### Health Check
GET /v1/healthcheck

##### Success response
HTTP Status Code: 200
```JSON
{
  "status": "ok"
}
```
