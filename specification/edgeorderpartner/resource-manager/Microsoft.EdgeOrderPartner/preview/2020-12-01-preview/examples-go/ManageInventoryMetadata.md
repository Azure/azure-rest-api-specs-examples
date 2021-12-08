Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fedgeorderpartner%2Farmedgeorderpartner%2Fv0.1.0/sdk/resourcemanager/edgeorderpartner/armedgeorderpartner/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/edgeorderpartner/resource-manager/Microsoft.EdgeOrderPartner/preview/2020-12-01-preview/examples/ManageInventoryMetadata.json
func ExampleEdgeOrderPartnerAPISClient_BeginManageInventoryMetadata() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armedgeorderpartner.NewEdgeOrderPartnerAPISClient("<subscription-id>", cred, nil)
	poller, err := client.BeginManageInventoryMetadata(ctx,
		"<family-identifier>",
		"<location>",
		"<serial-number>",
		armedgeorderpartner.ManageInventoryMetadataRequest{
			ConfigurationOnDevice: &armedgeorderpartner.ConfigurationOnDevice{
				ConfigurationIdentifier: to.StringPtr("<configuration-identifier>"),
			},
			InventoryMetadata: to.StringPtr("<inventory-metadata>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
```
