generate:
	protoc pb/usersvc/v1/usersvc.proto --go_out=plugins=grpc:.

#mock:
#	mockery --name=UserSvcClient --output=./mocks/