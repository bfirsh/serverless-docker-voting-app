
default: run

run: build
	docker-compose up --build

build:
	docker-compose -f docker-compose.build.yml build
