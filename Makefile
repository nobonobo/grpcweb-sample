
PID:=${shell pwd}/server.pid

start: $(PID)
	go run .

$(PID):
	cd frontend && { wasmserve & echo $$! > $@; }

stop:
	if [ -e $(PID) ]; then kill `cat $(PID)` && rm $(PID); fi

generate:
	go generate ./frontend

cert:
	go run $(shell go env GOROOT)/src/crypto/tls/generate_cert.go -ecdsa-curve P256 -host localhost -ca