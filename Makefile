build-images:
	docker build -t gateway-zero-ex/stream:latest -f stream/Dockerfile .

run-containers:
	docker compose up --build

run-stream:
	go run ./stream/stream.go -f ./stream/etc/stream-api.yaml 