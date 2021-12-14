generate:
	protoc pb/registrationsvc/v1/registrationsvc.proto --go_out=plugins=grpc:.