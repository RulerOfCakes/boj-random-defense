#!make
include .env

build:
	go build main.go

run:
	go run main.go -t=${BOT_TOKEN} -g=${GUILD_ID} -r=${REMOVE_COMMANDS}