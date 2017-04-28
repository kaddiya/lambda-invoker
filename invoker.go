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
func (i *AWSLambdaInvoker) InvokeLambda(request interface{}) (response map[string]interface{}, err error) {

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
	lambdaParams.FunctionName = aws.String(i.LambdaConfig.AWSLambdaFunctionName)
	lambdaParams.InvocationType = aws.String(i.LambdaConfig.AWSLamdaInvocationType)
	encodedPayload, marshalError := json.Marshal(request)

	if marshalError != nil {
		fmt.Println(marshalError)
		return nil, errors.New("There was an error in marshalling the payload for the invocation.Please check whether the object is properly formed")
	}

	lambdaParams.Payload = encodedPayload
	lambdaClient := lambda.New(config)
	lambdaHandler, err := lambdaClient.Invoke(lambdaParams)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("A problem was encountered in invoking the lambda")
	}

	payloadString := string(lambdaHandler.Payload)
	var result map[string]interface{}
	unmrshErr := json.Unmarshal([]byte(payloadString), &result)
	if unmrshErr != nil {
		fmt.Println(unmrshErr)
		return nil, errors.New("There was an error in parsing the lamdba Response")
	}
	return result, nil
}
