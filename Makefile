build:
	docker-compose build tgbot

run: build
	docker-compose up tgbot