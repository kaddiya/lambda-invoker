package models

// LambdaRequest encapsulates the request to Lambda
type LambdaRequest interface {
}

// LambdaResponse encapsulates the response recieved from Lambda
type LambdaResponse interface {
}

// LambdaInvoker will expose and API to invoke Lambda using the official go-aws-sdk
type LambdaInvoker interface {
	InvokeLambda(LambdaRequest) (LambdaResponse, error)
}
