package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fa469a1157c33837a46c9bcd524527e94125189a/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Subscription_TestQuery.json
func ExampleSubscriptionsClient_BeginTestQuery() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstreamanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSubscriptionsClient().BeginTestQuery(ctx, "West US", armstreamanalytics.TestQuery{
		Diagnostics: &armstreamanalytics.TestQueryDiagnostics{
			Path:     to.Ptr("/pathto/subdirectory"),
			WriteURI: to.Ptr("http://myoutput.com"),
		},
		StreamingJob: &armstreamanalytics.StreamingJob{
			Location: to.Ptr("West US"),
			Tags: map[string]*string{
				"key1":      to.Ptr("value1"),
				"key3":      to.Ptr("value3"),
				"randomKey": to.Ptr("randomValue"),
			},
			Properties: &armstreamanalytics.StreamingJobProperties{
				CompatibilityLevel:                 to.Ptr(armstreamanalytics.CompatibilityLevelOne0),
				DataLocale:                         to.Ptr("en-US"),
				EventsLateArrivalMaxDelayInSeconds: to.Ptr[int32](5),
				EventsOutOfOrderMaxDelayInSeconds:  to.Ptr[int32](0),
				EventsOutOfOrderPolicy:             to.Ptr(armstreamanalytics.EventsOutOfOrderPolicyDrop),
				Functions:                          []*armstreamanalytics.Function{},
				Inputs: []*armstreamanalytics.Input{
					{
						Name: to.Ptr("inputtest"),
						Properties: &armstreamanalytics.StreamInputProperties{
							Type: to.Ptr("Stream"),
							Serialization: &armstreamanalytics.JSONSerialization{
								Type: to.Ptr(armstreamanalytics.EventSerializationTypeJSON),
								Properties: &armstreamanalytics.JSONSerializationProperties{
									Encoding: to.Ptr(armstreamanalytics.EncodingUTF8),
								},
							},
							Datasource: &armstreamanalytics.RawStreamInputDataSource{
								Type: to.Ptr("Raw"),
								Properties: &armstreamanalytics.RawInputDatasourceProperties{
									PayloadURI: to.Ptr("http://myinput.com"),
								},
							},
						},
					}},
				OutputErrorPolicy: to.Ptr(armstreamanalytics.OutputErrorPolicyDrop),
				Outputs: []*armstreamanalytics.Output{
					{
						Name: to.Ptr("outputtest"),
						Properties: &armstreamanalytics.OutputProperties{
							Datasource: &armstreamanalytics.RawOutputDatasource{
								Type: to.Ptr("Raw"),
								Properties: &armstreamanalytics.RawOutputDatasourceProperties{
									PayloadURI: to.Ptr("http://myoutput.com"),
								},
							},
							Serialization: &armstreamanalytics.JSONSerialization{
								Type: to.Ptr(armstreamanalytics.EventSerializationTypeJSON),
							},
						},
					}},
				SKU: &armstreamanalytics.SKU{
					Name: to.Ptr(armstreamanalytics.SKUNameStandard),
				},
				Transformation: &armstreamanalytics.Transformation{
					Name: to.Ptr("transformationtest"),
					Properties: &armstreamanalytics.TransformationProperties{
						Query:          to.Ptr("Select Id, Name from inputtest"),
						StreamingUnits: to.Ptr[int32](1),
					},
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.QueryTestingResult = armstreamanalytics.QueryTestingResult{
	// 	OutputURI: to.Ptr("http://myoutput.com"),
	// 	Status: to.Ptr(armstreamanalytics.QueryTestingResultStatusSuccess),
	// }
}
