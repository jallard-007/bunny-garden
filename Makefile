.PHONY: files api

all: files api

files:
	cd frontend && npm run build
	go build ./cmd/files

api:
	go build ./cmd/api
