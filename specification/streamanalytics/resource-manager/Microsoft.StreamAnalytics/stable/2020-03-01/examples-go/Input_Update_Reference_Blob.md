Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstreamanalytics%2Farmstreamanalytics%2Fv0.5.0/sdk/resourcemanager/streamanalytics/armstreamanalytics/README.md) on how to add the SDK to your project and authenticate.

```go
package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Input_Update_Reference_Blob.json
func ExampleInputsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armstreamanalytics.NewInputsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<job-name>",
		"<input-name>",
		armstreamanalytics.Input{
			Properties: &armstreamanalytics.ReferenceInputProperties{
				Type: to.Ptr("<type>"),
				Serialization: &armstreamanalytics.CSVSerialization{
					Type: to.Ptr(armstreamanalytics.EventSerializationTypeCSV),
					Properties: &armstreamanalytics.CSVSerializationProperties{
						Encoding:       to.Ptr(armstreamanalytics.EncodingUTF8),
						FieldDelimiter: to.Ptr("<field-delimiter>"),
					},
				},
				Datasource: &armstreamanalytics.BlobReferenceInputDataSource{
					Type: to.Ptr("<type>"),
					Properties: &armstreamanalytics.BlobReferenceInputDataSourceProperties{
						Container: to.Ptr("<container>"),
					},
				},
			},
		},
		&armstreamanalytics.InputsClientUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
