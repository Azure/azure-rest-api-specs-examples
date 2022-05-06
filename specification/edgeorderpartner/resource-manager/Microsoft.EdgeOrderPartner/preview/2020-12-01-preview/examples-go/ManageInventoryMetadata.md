Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fedgeorderpartner%2Farmedgeorderpartner%2Fv0.4.0/sdk/resourcemanager/edgeorderpartner/armedgeorderpartner/README.md) on how to add the SDK to your project and authenticate.

```go
package armedgeorderpartner_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgeorderpartner/armedgeorderpartner"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/edgeorderpartner/resource-manager/Microsoft.EdgeOrderPartner/preview/2020-12-01-preview/examples/ManageInventoryMetadata.json
func ExampleAPISClient_BeginManageInventoryMetadata() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armedgeorderpartner.NewAPISClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginManageInventoryMetadata(ctx,
		"<family-identifier>",
		"<location>",
		"<serial-number>",
		armedgeorderpartner.ManageInventoryMetadataRequest{
			ConfigurationOnDevice: &armedgeorderpartner.ConfigurationOnDevice{
				ConfigurationIdentifier: to.Ptr("<configuration-identifier>"),
			},
			InventoryMetadata: to.Ptr("<inventory-metadata>"),
		},
		&armedgeorderpartner.APISClientBeginManageInventoryMetadataOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}
```
