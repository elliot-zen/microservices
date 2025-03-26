# Grpc microservices in go;

## Grpcurl

```sh 
# list service
grpcurl -plaintext localhost:3000 list
grpcurl -plaintext localhost:3000 describe <Service>

# call a service
grpcurl \
    -plaintext \
    -d '{"user_id": 123, "order_items": [{"product_code": "prod", "quantity": 4, "unit_price": 12}]}' \
    localhost:3000 \
    Order/Create
```
