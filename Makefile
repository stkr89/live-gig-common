generate:
	protoc pb/usersvc/v1/usersvc.proto --go_out=plugins=grpc:.

mock:
	mocker --dst pb/usersvc/v1/mock/usersvc_mock.go --pkg mock pb/usersvc/v1/usersvc.pb.go UserSvcClient
	mocker --dst pb/registrationsvc/v1/mock/registrationsvc_mock.go --pkg mock pb/registrationsvc/v1/registrationsvc.pb.go RegistrationSvcClient