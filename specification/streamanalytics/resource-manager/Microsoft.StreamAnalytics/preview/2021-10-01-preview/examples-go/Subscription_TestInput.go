package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fa469a1157c33837a46c9bcd524527e94125189a/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Subscription_TestInput.json
func ExampleSubscriptionsClient_BeginTestInput() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstreamanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSubscriptionsClient().BeginTestInput(ctx, "West US", armstreamanalytics.TestInput{
		Input: &armstreamanalytics.Input{
			Properties: &armstreamanalytics.StreamInputProperties{
				Type: to.Ptr("Stream"),
				Serialization: &armstreamanalytics.CSVSerialization{
					Type: to.Ptr(armstreamanalytics.EventSerializationTypeCSV),
					Properties: &armstreamanalytics.CSVSerializationProperties{
						Encoding:       to.Ptr(armstreamanalytics.EncodingUTF8),
						FieldDelimiter: to.Ptr(","),
					},
				},
				Datasource: &armstreamanalytics.BlobStreamInputDataSource{
					Type: to.Ptr("Microsoft.Storage/Blob"),
					Properties: &armstreamanalytics.BlobStreamInputDataSourceProperties{
						Container:   to.Ptr("state"),
						DateFormat:  to.Ptr("yyyy/MM/dd"),
						PathPattern: to.Ptr("{date}/{time}"),
						StorageAccounts: []*armstreamanalytics.StorageAccount{
							{
								AccountKey:  to.Ptr("someAccountKey=="),
								AccountName: to.Ptr("someAccountName"),
							}},
						TimeFormat:           to.Ptr("HH"),
						SourcePartitionCount: to.Ptr[int32](16),
					},
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
