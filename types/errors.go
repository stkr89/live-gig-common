package types

import "errors"

var (
	InvalidRequestBody  = errors.New("invalid request body")
	AWSUserSignupFailed = errors.New("unable to signup user in AWS Cognito")
	AddUserToDBFailed   = errors.New("unable to add user to db")
)
