Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstreamanalytics%2Farmstreamanalytics%2Fv0.3.0/sdk/resourcemanager/streamanalytics/armstreamanalytics/README.md) on how to add the SDK to your project and authenticate.

```go
package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics"
)

// x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Output_Update_DocumentDB.json
func ExampleOutputsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstreamanalytics.NewOutputsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<job-name>",
		"<output-name>",
		armstreamanalytics.Output{
			Properties: &armstreamanalytics.OutputProperties{
				Datasource: &armstreamanalytics.DocumentDbOutputDataSource{
					Type: to.StringPtr("<type>"),
					Properties: &armstreamanalytics.DocumentDbOutputDataSourceProperties{
						PartitionKey: to.StringPtr("<partition-key>"),
					},
				},
			},
		},
		&armstreamanalytics.OutputsClientUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.OutputsClientUpdateResult)
}
```
