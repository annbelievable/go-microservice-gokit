.PHONY: protos

protos:
	 protoc -I protos/ protos/contact.proto --go_out=plugins=grpc:protos/ 