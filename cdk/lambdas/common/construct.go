package common

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
)

type GoLambdaConstructProps struct {
	awslambda.FunctionProps
}

func NewGoLambdaConstruct(scope constructs.Construct, id *string, props *GoLambdaConstructProps) awslambda.Function {
	return awslambda.NewFunction(scope, id, &props.FunctionProps)
}
