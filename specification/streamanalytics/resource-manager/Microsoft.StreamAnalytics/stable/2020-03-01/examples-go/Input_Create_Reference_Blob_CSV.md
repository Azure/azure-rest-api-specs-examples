Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstreamanalytics%2Farmstreamanalytics%2Fv0.3.1/sdk/resourcemanager/streamanalytics/armstreamanalytics/README.md) on how to add the SDK to your project and authenticate.

```go
package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics"
)

// x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Input_Create_Reference_Blob_CSV.json
func ExampleInputsClient_CreateOrReplace() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstreamanalytics.NewInputsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrReplace(ctx,
		"<resource-group-name>",
		"<job-name>",
		"<input-name>",
		armstreamanalytics.Input{
			Properties: &armstreamanalytics.ReferenceInputProperties{
				Type: to.StringPtr("<type>"),
				Serialization: &armstreamanalytics.CSVSerialization{
					Type: armstreamanalytics.EventSerializationType("Csv").ToPtr(),
					Properties: &armstreamanalytics.CSVSerializationProperties{
						Encoding:       armstreamanalytics.Encoding("UTF8").ToPtr(),
						FieldDelimiter: to.StringPtr("<field-delimiter>"),
					},
				},
				Datasource: &armstreamanalytics.BlobReferenceInputDataSource{
					Type: to.StringPtr("<type>"),
					Properties: &armstreamanalytics.BlobReferenceInputDataSourceProperties{
						Container:   to.StringPtr("<container>"),
						DateFormat:  to.StringPtr("<date-format>"),
						PathPattern: to.StringPtr("<path-pattern>"),
						StorageAccounts: []*armstreamanalytics.StorageAccount{
							{
								AccountKey:  to.StringPtr("<account-key>"),
								AccountName: to.StringPtr("<account-name>"),
							}},
						TimeFormat: to.StringPtr("<time-format>"),
					},
				},
			},
		},
		&armstreamanalytics.InputsClientCreateOrReplaceOptions{IfMatch: nil,
			IfNoneMatch: nil,
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.InputsClientCreateOrReplaceResult)
}
```
