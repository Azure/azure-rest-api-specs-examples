```go
package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Input_Create_Reference_Blob_CSV.json
func ExampleInputsClient_CreateOrReplace() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstreamanalytics.NewInputsClient("56b5e0a9-b645-407d-99b0-c64f86013e3d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrReplace(ctx,
		"sjrg8440",
		"sj9597",
		"input7225",
		armstreamanalytics.Input{
			Properties: &armstreamanalytics.ReferenceInputProperties{
				Type: to.Ptr("Reference"),
				Serialization: &armstreamanalytics.CSVSerialization{
					Type: to.Ptr(armstreamanalytics.EventSerializationTypeCSV),
					Properties: &armstreamanalytics.CSVSerializationProperties{
						Encoding:       to.Ptr(armstreamanalytics.EncodingUTF8),
						FieldDelimiter: to.Ptr(","),
					},
				},
				Datasource: &armstreamanalytics.BlobReferenceInputDataSource{
					Type: to.Ptr("Microsoft.Storage/Blob"),
					Properties: &armstreamanalytics.BlobReferenceInputDataSourceProperties{
						Container:   to.Ptr("state"),
						DateFormat:  to.Ptr("yyyy/MM/dd"),
						PathPattern: to.Ptr("{date}/{time}"),
						StorageAccounts: []*armstreamanalytics.StorageAccount{
							{
								AccountKey:  to.Ptr("someAccountKey=="),
								AccountName: to.Ptr("someAccountName"),
							}},
						TimeFormat: to.Ptr("HH"),
					},
				},
			},
		},
		&armstreamanalytics.InputsClientCreateOrReplaceOptions{IfMatch: nil,
			IfNoneMatch: nil,
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstreamanalytics%2Farmstreamanalytics%2Fv1.0.0/sdk/resourcemanager/streamanalytics/armstreamanalytics/README.md) on how to add the SDK to your project and authenticate.
