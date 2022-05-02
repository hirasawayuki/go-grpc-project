# go-grpc-project
Learn the basics of　gRPC with Go.

## usage

```
$ docker-compose build .
$ docker-compose up -d db
$ docker-compose up -d
```

### Endpoint

**POST /auth/register/**

signup user.
params: email, password
```
$ curl --request POST \
  --url http://localhost:3000/auth/register/ \
  --header 'Content-Type: application/json' \
  --data '{
 "email": "example@example.com",
 "password": "password"
}'
```

**POST /auth/login/**

user login.
params: email, password
```
curl --request POST \
  --url http://localhost:3000/auth/login/ \
  --header 'Content-Type: application/json' \
  --data '{
 "email": "example@example.com",
 "password": "password"
}'

response
{
 "status": 200,
 "token": "jwt token"
}
```

**POST /product/**

create product.
params: name, stock, price

```
curl --request POST \
  --url http://localhost:3000/product \
  --header 'Authorization: Bearer {jwt token}' \
  --header 'Content-Type: application/json' \
  --data '{
 "name": "Product A",
 "stock": 5,
 "price": 15
}'
```
**GET /product/:id**

Find one product.
params: id
```
curl --request GET \
  --url http://localhost:3000/product/{id} \
  --header 'Authorization: Bearer {jwt token}'
```

**POST /order/**

create order
params: product_id, quantity
```
curl --request POST \
  --url http://localhost:3000/order \
  --header 'Authorization: Bearer {jwt token}' \
  --header 'Content-Type: application/json' \
  --data '{
 "productId": 1,
 "quantity": 1
}'
```
