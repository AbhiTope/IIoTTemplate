curl -X POST http://localhost:8080/login -d '{"userName":"manu", "password":"123"}' -H "Content-Type: application/json"
echo ''
curl -X POST http://localhost:8080/login -d '{"userName":"user2", "password":"456"}' -H "Content-Type: application/json"

