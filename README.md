# lambda-invoker
A general utility package to invoke aws lambda from your Go code

**Usage**  
in the calling code
```
//Declare a struct ActualAWSConfigProvider which will provide the implementation for BaseAWSConfigProvider
type EnvironmentVarsAWSConfigProvider struct{

}

// Make it implement the BaseAWSConfigProvider interface by supplying implementation of GetBaseAWSConfig() which will provide  // the credentials by sourcing it from somewhere like os.Getenv()
func(a *EnvironmentVarsAWSConfigProvider)GetBaseAWSConfig()(BaseAWSConfig, error){
  return &BaseAWSConfig{
    AWSRegion: os.Getenv("GET_YOUR_REGION"),
    AWSAccessKey:os.Getenv("ACCESS_KEY"),
    AWSSecretKey:os.Getenv("SECRET_KEY"),
  },
}


//create the instance of the AWSLambdaInvoker
invoker := &AWSLambdaInvoker{
  LambdaConfig: &AWSLambdaConfig{
    AWSLambdaFunctionName:"SomeFunctionInLambda",
    AWSLamdaInvocationType:"RequestResonse",
  },
  AWSConfigProvider: &EnvironmentVarsAWSConfigProvider{}
}

// CustomLambdaRequest will provide the encapsulation for the request
type CustomLambdaRequest struct{
  //model the request
}
//instantiate the request
req := CustomLambdaRequest{}

//invoke the method
res := invoker.InvokeLambda(req)

//res will have the invocation response as a map of string and arbitary values for your perusal

```


**Support**  
1.Currently supports only `RequestResonse` invocation type.  
2.Others coming soon.
