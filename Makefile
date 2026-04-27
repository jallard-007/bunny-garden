.PHONY: files api start-files start-api

all: files api

files:
	cd frontend && npm run build
	go build ./cmd/bg-files

api:
	go build ./cmd/bg-api

start-files:
	pkill -fx "./bg-files --port 8061" || true
	nohup ./bg-files --port 8061 &> files.log &

start-api:
	pkill -fx "./bg-api serve --http 127.0.0.1:8062" || true
	nohup ./bg-api serve --http 127.0.0.1:8062 &> api.log &
