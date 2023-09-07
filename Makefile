#!make

# DOCKER
docker-up: 
	docker-compose up -d

docker-down: 
	docker-compose down

docker-build:
	docker build --tag dealls .