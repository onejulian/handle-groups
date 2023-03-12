## Implementación de una función lambda + api gateway
### Node
Abra una terminal bash y ejecute el siguiente comando:

```bash
curl --location 'https://5tcjj7217h.execute-api.us-east-2.amazonaws.com/default/handle-groups' \
--header 'Content-Type: application/json' \
--data '{
  "groups": "1,2,1,1,1,2,1,3"
}'
```