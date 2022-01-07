generate:
	protoc pb/usersvc/v1/usersvc.proto --go_out=plugins=grpc:.
	protoc pb/authsvc/v1/authsvc.proto --go_out=plugins=grpc:.
	protoc pb/orgsvc/v1/orgsvc.proto --go_out=plugins=grpc:.
	protoc pb/ticketsvc/v1/ticketsvc.proto --go_out=plugins=grpc:.

mock:
	mocker --dst pb/usersvc/v1/mock/usersvc_mock.go --pkg mock pb/usersvc/v1/usersvc.pb.go UserSvcClient
	mocker --dst pb/authsvc/v1/mock/authsvc_mock.go --pkg mock pb/authsvc/v1/authsvc.pb.go AuthSvcClient
	mocker --dst pb/orgsvc/v1/mock/orgsvc_mock.go --pkg mock pb/orgsvc/v1/orgsvc.pb.go OrgSvcClient
	mocker --dst pb/ticketsvc/v1/mock/ticketsvc_mock.go --pkg mock pb/ticketsvc/v1/ticketsvc.pb.go TicketSvcClient