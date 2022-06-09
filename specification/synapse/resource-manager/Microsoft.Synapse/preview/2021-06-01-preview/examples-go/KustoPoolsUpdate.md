```go
package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolsUpdate.json
func ExampleKustoPoolsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewKustoPoolsClient("12345678-1234-1234-1234-123456789098", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"synapseWorkspaceName",
		"kustorptest",
		"kustoclusterrptest4",
		armsynapse.KustoPoolUpdate{
			Properties: &armsynapse.KustoPoolProperties{
				EnablePurge:           to.Ptr(true),
				EnableStreamingIngest: to.Ptr(true),
				WorkspaceUID:          to.Ptr("11111111-2222-3333-444444444444"),
			},
			SKU: &armsynapse.AzureSKU{
				Name:     to.Ptr(armsynapse.SKUNameStorageOptimized),
				Capacity: to.Ptr[int32](2),
				Size:     to.Ptr(armsynapse.SKUSizeMedium),
			},
		},
		&armsynapse.KustoPoolsClientBeginUpdateOptions{IfMatch: nil})
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsynapse%2Farmsynapse%2Fv0.5.0/sdk/resourcemanager/synapse/armsynapse/README.md) on how to add the SDK to your project and authenticate.
