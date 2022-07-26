# List API
## Auth
#### Login
```
curl --request POST \
  --url http://localhost:8081/login \
  --header 'content-type: application/json' \
  --data '{
	"email": "email01",
	"password": "password01"
}'
```
#### Registration for buyer
```
curl --request POST \
  --url http://localhost:8081/register \
  --header 'content-type: application/json' \
  --data '{
	"email": "email02",
	"password": "password02",
	"name": "name02",
	"address": "address02",
	"role": "buyer"
}'
```

#### registration for seller
```
curl --request POST \
  --url http://localhost:8081/register \
  --header 'content-type: application/json' \
  --data '{
	"email": "email01",
	"password": "password01",
	"name": "name01",
	"address": "address01",
	"role": "seller"
}'
```

## Product

#### Create Product
```
curl --request POST \
  --url http://localhost:8081/v1/createProduct \
  --header 'authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im5hbWUwMSIsInVzZXJfaWQiOjEsInJvbGUiOiJzZWxsZXIiLCJleHAiOjE2NTg4MjgzMzh9.EE5l9LHtusXn2XQSFHccmolZPzWS4p7ccUfDTiS-ioE' \
  --header 'content-type: application/json' \
  --data '{
	"name": "name 01",
	"description": "decription 01",
	"price": 10222
}'
```

#### Get list product
```
curl --request POST \
  --url http://localhost:8081/v1/getProducts \
  --header 'authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im5hbWUwMSIsInVzZXJfaWQiOjEsInJvbGUiOiJzZWxsZXIiLCJleHAiOjE2NTg4MjgzMzh9.EE5l9LHtusXn2XQSFHccmolZPzWS4p7ccUfDTiS-ioE' \
  --header 'content-type: application/json' \
  --data '{
	"seller_id": 1,
	"page": 1,
	"per_page": 100
}'
```

## Order 
#### Create order

```
curl --request POST \
  --url http://localhost:8081/v1/createOrder \
  --header 'authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im5hbWUwMiIsInVzZXJfaWQiOjIsInJvbGUiOiJidXllciIsImV4cCI6MTY1ODgzMjA4M30.KxS4-K8drNnEHfJ7ZEjJSMgsNqnTUkE_Lz-zfM7cjNg' \
  --header 'content-type: application/json' \
  --data '{
	"order_items":[{
			"product_id": 1,
			"quantity":1
	},
	{
			"product_id": 2,
			"quantity":1
	},
	{
			"product_id": 3,
			"quantity":2
	}
	] 
}'
```

#### Get list order 

```
curl --request POST \
  --url http://localhost:8081/v1/getOrders \
  --header 'authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im5hbWUwMiIsInVzZXJfaWQiOjIsInJvbGUiOiJidXllciIsImV4cCI6MTY1ODgzMjA4M30.KxS4-K8drNnEHfJ7ZEjJSMgsNqnTUkE_Lz-zfM7cjNg' \
  --header 'content-type: application/json' \
  --data '{
	"page": 1,
	"per_page":100
}'
```

#### Accept order

```
curl --request POST \
  --url http://localhost:8081/v1/acceptOrder \
  --header 'authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im5hbWUwMSIsInVzZXJfaWQiOjEsInJvbGUiOiJzZWxsZXIiLCJleHAiOjE2NTg4MzM2NzZ9.fZsEa7bC3dstTLKnN8a9UsIjVQEN-FjQek8HYMNvEYM' \
  --header 'content-type: application/json' \
  --data '{
	"order_id": 2
}'
```
