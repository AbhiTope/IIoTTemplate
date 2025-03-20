echo ''

curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{
    "userName": "john_doe",
    "firstName": "John",
    "lastName": "Doe",
    "email": "john.doe@example.com",
    "role": "user",
    "password": "securepassword123"
  }'

