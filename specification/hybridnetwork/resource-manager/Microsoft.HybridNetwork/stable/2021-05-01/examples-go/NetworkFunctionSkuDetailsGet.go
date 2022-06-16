package armhybridnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork"
)

// x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2021-05-01/examples/NetworkFunctionSkuDetailsGet.json
func ExampleNetworkFunctionVendorSKUsClient_ListBySKU() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhybridnetwork.NewNetworkFunctionVendorSKUsClient("<subscription-id>", cred, nil)
	pager := client.ListBySKU("<vendor-name>",
		"<vendor-skuname>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
	}
}
