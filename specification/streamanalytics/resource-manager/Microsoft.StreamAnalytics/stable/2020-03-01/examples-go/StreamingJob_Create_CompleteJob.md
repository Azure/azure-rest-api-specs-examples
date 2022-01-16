Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstreamanalytics%2Farmstreamanalytics%2Fv0.3.0/sdk/resourcemanager/streamanalytics/armstreamanalytics/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/StreamingJob_Create_CompleteJob.json
func ExampleStreamingJobsClient_BeginCreateOrReplace() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstreamanalytics.NewStreamingJobsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrReplace(ctx,
		"<resource-group-name>",
		"<job-name>",
		armstreamanalytics.StreamingJob{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"key1":      to.StringPtr("value1"),
				"key3":      to.StringPtr("value3"),
				"randomKey": to.StringPtr("randomValue"),
			},
			Properties: &armstreamanalytics.StreamingJobProperties{
				CompatibilityLevel:                 armstreamanalytics.CompatibilityLevel("1.0").ToPtr(),
				DataLocale:                         to.StringPtr("<data-locale>"),
				EventsLateArrivalMaxDelayInSeconds: to.Int32Ptr(5),
				EventsOutOfOrderMaxDelayInSeconds:  to.Int32Ptr(0),
				EventsOutOfOrderPolicy:             armstreamanalytics.EventsOutOfOrderPolicy("Drop").ToPtr(),
				Functions:                          []*armstreamanalytics.Function{},
				Inputs: []*armstreamanalytics.Input{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armstreamanalytics.StreamInputProperties{
							Type: to.StringPtr("<type>"),
							Serialization: &armstreamanalytics.JSONSerialization{
								Type: armstreamanalytics.EventSerializationType("Json").ToPtr(),
								Properties: &armstreamanalytics.JSONSerializationProperties{
									Encoding: armstreamanalytics.Encoding("UTF8").ToPtr(),
								},
							},
							Datasource: &armstreamanalytics.BlobStreamInputDataSource{
								Type: to.StringPtr("<type>"),
								Properties: &armstreamanalytics.BlobStreamInputDataSourceProperties{
									Container:   to.StringPtr("<container>"),
									PathPattern: to.StringPtr("<path-pattern>"),
									StorageAccounts: []*armstreamanalytics.StorageAccount{
										{
											AccountKey:  to.StringPtr("<account-key>"),
											AccountName: to.StringPtr("<account-name>"),
										}},
								},
							},
						},
					}},
				OutputErrorPolicy: armstreamanalytics.OutputErrorPolicy("Drop").ToPtr(),
				Outputs: []*armstreamanalytics.Output{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armstreamanalytics.OutputProperties{
							Datasource: &armstreamanalytics.AzureSQLDatabaseOutputDataSource{
								Type: to.StringPtr("<type>"),
								Properties: &armstreamanalytics.AzureSQLDatabaseOutputDataSourceProperties{
									Database: to.StringPtr("<database>"),
									Password: to.StringPtr("<password>"),
									Server:   to.StringPtr("<server>"),
									Table:    to.StringPtr("<table>"),
									User:     to.StringPtr("<user>"),
								},
							},
						},
					}},
				SKU: &armstreamanalytics.SKU{
					Name: armstreamanalytics.SKUName("Standard").ToPtr(),
				},
				Transformation: &armstreamanalytics.Transformation{
					Name: to.StringPtr("<name>"),
					Properties: &armstreamanalytics.TransformationProperties{
						Query:          to.StringPtr("<query>"),
						StreamingUnits: to.Int32Ptr(1),
					},
				},
			},
		},
		&armstreamanalytics.StreamingJobsClientBeginCreateOrReplaceOptions{IfMatch: nil,
			IfNoneMatch: nil,
		})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.StreamingJobsClientCreateOrReplaceResult)
}
```
