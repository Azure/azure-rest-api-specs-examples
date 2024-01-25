package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fa469a1157c33837a46c9bcd524527e94125189a/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Subscription_CompileQuery.json
func ExampleSubscriptionsClient_CompileQuery() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstreamanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSubscriptionsClient().CompileQuery(ctx, "West US", armstreamanalytics.CompileQuery{
		CompatibilityLevel: to.Ptr(armstreamanalytics.CompatibilityLevelOne2),
		Functions: []*armstreamanalytics.QueryFunction{
			{
				Name:        to.Ptr("function1"),
				Type:        to.Ptr("Scalar"),
				BindingType: to.Ptr("Microsoft.StreamAnalytics/JavascriptUdf"),
				Inputs: []*armstreamanalytics.FunctionInput{
					{
						DataType: to.Ptr("any"),
					}},
				Output: &armstreamanalytics.FunctionOutput{
					DataType: to.Ptr("bigint"),
				},
			}},
		Inputs: []*armstreamanalytics.QueryInput{
			{
				Name: to.Ptr("input1"),
				Type: to.Ptr("Stream"),
			}},
		JobType: to.Ptr(armstreamanalytics.JobTypeCloud),
		Query:   to.Ptr("SELECT\r\n    *\r\nINTO\r\n    [output1]\r\nFROM\r\n    [input1]"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.QueryCompilationResult = armstreamanalytics.QueryCompilationResult{
	// 	Errors: []*armstreamanalytics.QueryCompilationError{
	// 		{
	// 			IsGlobal: to.Ptr(true),
	// 			Message: to.Ptr("Query failed to compile."),
	// 	}},
	// 	Functions: []*string{
	// 		to.Ptr("transformationtest")},
	// 		Inputs: []*string{
	// 			to.Ptr("inputtest")},
	// 			Outputs: []*string{
	// 				to.Ptr("outputtest")},
	// 			}
}
