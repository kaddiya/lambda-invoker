package lambdainvoker

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

// BaseAWSConfig will hold the values that are generally required by all things for AWS
type BaseAWSConfig struct {
	AWSRegion    string
	AWSAccessKey string
	AWSSecretKey string
}

// AWSLambdaConfig will encapsulate the config required for invoking a lambda
type AWSLambdaConfig struct {
	AWSConfig              BaseAWSConfig
	AWSLambdaFunctionName  string
	AWSLamdaInvocationType string
}

type AWSLambdaInvoker struct {
	Config AWSLambdaConfig
}
