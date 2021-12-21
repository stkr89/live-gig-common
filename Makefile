generate:
	protoc pb/signupsvc/v1/signupsvc.proto --go_out=plugins=grpc:.

mock:
	mocker --dst pb/usersvc/v1/mock/usersvc_mock.go --pkg mock pb/usersvc/v1/usersvc.pb.go UserSvcClient
	mocker --dst pb/signupsvc/v1/mock/signupsvc_mock.go --pkg mock pb/signupsvc/v1/signupsvc.pb.go SignUpSvcClient