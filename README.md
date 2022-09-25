Order Service
===

This GO project serves as a microservice for [eCommerce](https://github.com/users/ethmore/projects/4) project.


## Service task:

- Request address and cart info from `auth-and-db-service` create and send order object to [auth-and-db-service](https://github.com/ethmore/auth-and-db-service)



# Installation

Ensure GO is installed on your system
```
go mod download
````

```
go run .
```

## Test
```
curl http://localhost:3009/test
```
### It should return:
```
StatusCode        : 200
StatusDescription : OK
Content           : {"message":"OK"}
```

## Example .env file
This file should be placed inside `dotEnv` folder
```
# Cors URLs
BFFURL = http://localhost:3001

# Request URLs
GETUSERADDRESSBYID = http://127.0.0.1:3002/getUserAddressById
GETCARTINFO = http://127.0.0.1:3002/getCartInfo
GETCARTPRODUCTS = http://127.0.0.1:3002/getCartProducts
INSERTORDER = http://127.0.0.1:3002/insertOrder
GETPRODUCTSSELLERS = http://127.0.0.1:3002/getProductsSellers
```