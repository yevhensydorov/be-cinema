# be-cinema

Golang backend server for cinema directory project

# Running in Docker

docker build --tag be-cinema-test .
docker run -p 8080:8080 --env-file .env be-cinema-test:latest
