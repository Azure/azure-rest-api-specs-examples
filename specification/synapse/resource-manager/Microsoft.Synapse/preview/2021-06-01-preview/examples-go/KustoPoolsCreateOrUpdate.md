Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsynapse%2Farmsynapse%2Fv0.2.1/sdk/resourcemanager/synapse/armsynapse/README.md) on how to add the SDK to your project and authenticate.

```go
package armsynapse_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolsCreateOrUpdate.json
func ExampleKustoPoolsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewKustoPoolsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<workspace-name>",
		"<resource-group-name>",
		"<kusto-pool-name>",
		armsynapse.KustoPool{
			Location: to.StringPtr("<location>"),
			Properties: &armsynapse.KustoPoolProperties{
				EnablePurge:           to.BoolPtr(true),
				EnableStreamingIngest: to.BoolPtr(true),
				WorkspaceUID:          to.StringPtr("<workspace-uid>"),
			},
			SKU: &armsynapse.AzureSKU{
				Name:     armsynapse.SKUName("Storage optimized").ToPtr(),
				Capacity: to.Int32Ptr(2),
				Size:     armsynapse.SKUSize("Medium").ToPtr(),
			},
		},
		&armsynapse.KustoPoolsClientBeginCreateOrUpdateOptions{IfMatch: nil,
			IfNoneMatch: nil,
		})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.KustoPoolsClientCreateOrUpdateResult)
}
```
