PROTO_DIR=proto
ENVOY_DIR=envoy
GEN_DIR=.
PROTOC_GEN_GO=$(shell which protoc)

all: clean generate

generate:
	@echo "Generating gRPC code..."
	@mkdir -p $(GEN_DIR)
	$(PROTOC_GEN_GO) --go_out=$(GEN_DIR) --go-grpc_out=$(GEN_DIR) $(PROTO_DIR)/noteapp.proto
	@echo "gRPC code generated successfully in $(GEN_DIR)"

clean:
	@echo "Cleaning generated files..."
	@rm -rf $(PROTO_DIR)/generated
	@echo "Cleaned successfully!"

start:
	go run cmd/main.go

createtopic:
	docker exec -it kafka kafka-topics --create --topic update-notes-topic --bootstrap-server kafka:29092 --partitions 1 --replication-factor 1

redisstart:
	docker run --name redis -p 6379:6379 -d redis:latest

envoyproto:
	protoc --descriptor_set_out=$(ENVOY_DIR)/proto.pb --include_imports --proto_path=$(ENVOY_DIR)/ $(ENVOY_DIR)/*.proto

envoystart:
	docker run -d --name envoy \
  	-p 9901:9901 -p 8080:8080 -p 9000:9000 \
  	-v $(pwd)/envoy/envoy.yaml:/etc/envoy/envoy.yaml \
  	-v $(pwd)/envoy/proto.pb:/etc/envoy/proto.pb \
  	envoyproxy/envoy:v1.28.0

envoystop:
	docker stop envoy && docker rm envoy


	





