package lambdainvoker

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

// InvokeLambda actually invokes the lambda
func (i *AWSLambdaInvoker) InvokeLambda(request LambdaRequest) (response LambdaResponse, err error) {

	baseCfg, cfgErr := i.AWSConfigProvider.GetBaseAWSConfig()
	if cfgErr != nil {
		return nil, errors.New("Could not build the base AWS Config")
	}

	creds := credentials.NewStaticCredentials(baseCfg.AWSAccessKey, baseCfg.AWSSecretKey, "")

	config := session.New(&aws.Config{
		Region:      aws.String(baseCfg.AWSRegion),
		Credentials: creds,
	})

	lambdaParams := &lambda.InvokeInput{}
	lambdaParams.FunctionName = aws.String(i.Config.AWSLambdaFunctionName)
	lambdaParams.InvocationType = aws.String(i.Config.AWSLamdaInvocationType)
	encodedPayload, marshalError := json.Marshal(request)

	if marshalError != nil {
		fmt.Println(marshalError)
		return nil, marshalError
	}

	lambdaParams.Payload = encodedPayload
	lambdaClient := lambda.New(config)
	lambdaHandler, err := lambdaClient.Invoke(lambdaParams)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	payloadString := string(lambdaHandler.Payload)
	fmt.Println(payloadString)
	return new(LambdaResponse), nil
}
