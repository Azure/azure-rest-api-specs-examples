Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstreamanalytics%2Farmstreamanalytics%2Fv0.5.0/sdk/resourcemanager/streamanalytics/armstreamanalytics/README.md) on how to add the SDK to your project and authenticate.

```go
package armstreamanalytics_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/StreamingJob_Create_CompleteJob.json
func ExampleStreamingJobsClient_BeginCreateOrReplace() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armstreamanalytics.NewStreamingJobsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrReplace(ctx,
		"<resource-group-name>",
		"<job-name>",
		armstreamanalytics.StreamingJob{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"key1":      to.Ptr("value1"),
				"key3":      to.Ptr("value3"),
				"randomKey": to.Ptr("randomValue"),
			},
			Properties: &armstreamanalytics.StreamingJobProperties{
				CompatibilityLevel:                 to.Ptr(armstreamanalytics.CompatibilityLevelOne0),
				DataLocale:                         to.Ptr("<data-locale>"),
				EventsLateArrivalMaxDelayInSeconds: to.Ptr[int32](5),
				EventsOutOfOrderMaxDelayInSeconds:  to.Ptr[int32](0),
				EventsOutOfOrderPolicy:             to.Ptr(armstreamanalytics.EventsOutOfOrderPolicyDrop),
				Functions:                          []*armstreamanalytics.Function{},
				Inputs: []*armstreamanalytics.Input{
					{
						Name: to.Ptr("<name>"),
						Properties: &armstreamanalytics.StreamInputProperties{
							Type: to.Ptr("<type>"),
							Serialization: &armstreamanalytics.JSONSerialization{
								Type: to.Ptr(armstreamanalytics.EventSerializationTypeJSON),
								Properties: &armstreamanalytics.JSONSerializationProperties{
									Encoding: to.Ptr(armstreamanalytics.EncodingUTF8),
								},
							},
							Datasource: &armstreamanalytics.BlobStreamInputDataSource{
								Type: to.Ptr("<type>"),
								Properties: &armstreamanalytics.BlobStreamInputDataSourceProperties{
									Container:   to.Ptr("<container>"),
									PathPattern: to.Ptr("<path-pattern>"),
									StorageAccounts: []*armstreamanalytics.StorageAccount{
										{
											AccountKey:  to.Ptr("<account-key>"),
											AccountName: to.Ptr("<account-name>"),
										}},
								},
							},
						},
					}},
				OutputErrorPolicy: to.Ptr(armstreamanalytics.OutputErrorPolicyDrop),
				Outputs: []*armstreamanalytics.Output{
					{
						Name: to.Ptr("<name>"),
						Properties: &armstreamanalytics.OutputProperties{
							Datasource: &armstreamanalytics.AzureSQLDatabaseOutputDataSource{
								Type: to.Ptr("<type>"),
								Properties: &armstreamanalytics.AzureSQLDatabaseOutputDataSourceProperties{
									Database: to.Ptr("<database>"),
									Password: to.Ptr("<password>"),
									Server:   to.Ptr("<server>"),
									Table:    to.Ptr("<table>"),
									User:     to.Ptr("<user>"),
								},
							},
						},
					}},
				SKU: &armstreamanalytics.SKU{
					Name: to.Ptr(armstreamanalytics.SKUNameStandard),
				},
				Transformation: &armstreamanalytics.Transformation{
					Name: to.Ptr("<name>"),
					Properties: &armstreamanalytics.TransformationProperties{
						Query:          to.Ptr("<query>"),
						StreamingUnits: to.Ptr[int32](1),
					},
				},
			},
		},
		&armstreamanalytics.StreamingJobsClientBeginCreateOrReplaceOptions{IfMatch: nil,
			IfNoneMatch: nil,
			ResumeToken: "",
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
