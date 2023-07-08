#!make
include envfile

build:
	go build main.go

run:
	go run main.go -t=${BOT_TOKEN}