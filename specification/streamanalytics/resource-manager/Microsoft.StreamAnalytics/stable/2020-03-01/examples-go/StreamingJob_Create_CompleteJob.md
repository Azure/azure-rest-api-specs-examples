```go
package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/StreamingJob_Create_CompleteJob.json
func ExampleStreamingJobsClient_BeginCreateOrReplace() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstreamanalytics.NewStreamingJobsClient("56b5e0a9-b645-407d-99b0-c64f86013e3d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrReplace(ctx,
		"sjrg3276",
		"sj7804",
		armstreamanalytics.StreamingJob{
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
							Datasource: &armstreamanalytics.BlobStreamInputDataSource{
								Type: to.Ptr("Microsoft.Storage/Blob"),
								Properties: &armstreamanalytics.BlobStreamInputDataSourceProperties{
									Container:   to.Ptr("containerName"),
									PathPattern: to.Ptr(""),
									StorageAccounts: []*armstreamanalytics.StorageAccount{
										{
											AccountKey:  to.Ptr("yourAccountKey=="),
											AccountName: to.Ptr("yourAccountName"),
										}},
								},
							},
						},
					}},
				OutputErrorPolicy: to.Ptr(armstreamanalytics.OutputErrorPolicyDrop),
				Outputs: []*armstreamanalytics.Output{
					{
						Name: to.Ptr("outputtest"),
						Properties: &armstreamanalytics.OutputProperties{
							Datasource: &armstreamanalytics.AzureSQLDatabaseOutputDataSource{
								Type: to.Ptr("Microsoft.Sql/Server/Database"),
								Properties: &armstreamanalytics.AzureSQLDatabaseOutputDataSourceProperties{
									Database: to.Ptr("databaseName"),
									Password: to.Ptr("userPassword"),
									Server:   to.Ptr("serverName"),
									Table:    to.Ptr("tableName"),
									User:     to.Ptr("<user>"),
								},
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
		&armstreamanalytics.StreamingJobsClientBeginCreateOrReplaceOptions{IfMatch: nil,
			IfNoneMatch: nil,
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstreamanalytics%2Farmstreamanalytics%2Fv1.0.0/sdk/resourcemanager/streamanalytics/armstreamanalytics/README.md) on how to add the SDK to your project and authenticate.
