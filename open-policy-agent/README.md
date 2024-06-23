
# https://medium.com/swlh/securing-dockerized-microservices-with-open-policy-agent-and-envoy-c128dfc764fe
```
// Access Denited
curl -X POST http://localhost:8080/authorize -d '{
  "user": "john",
  "action": "write",
  "resource": "mysql"
}' -H "Content-Type: application/json"

// Access Granted
curl -X POST http://localhost:8080/authorize -d '{
  "user": "john",
  "action": "write",
  "resource": "mysql"
}' -H "Content-Type: application/json"
```
