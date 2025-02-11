# Test API (Using curl or Postman)
## 📌 Create a User
curl -X POST http://localhost:8080/users \
     -H "Content-Type: application/json" \
     -d '{"name": "Alice", "age": 28, "address": "123 Main St", "salary": 50000}'

## 📌 Get All Users
curl http://localhost:8080/users

## 📌 Get a Specific User
curl http://localhost:8080/users/1

## 📌 Update a User
curl -X PUT http://localhost:8080/users/1 \
     -H "Content-Type: application/json" \
     -d '{"name": "Alice Updated", "age": 30, "address": "456 Elm St", "salary": 55000}'

## 📌 Delete a User
curl -X DELETE http://localhost:8080/users/1