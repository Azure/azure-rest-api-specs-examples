Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdnsresolver%2Farmdnsresolver%2Fv0.1.0/sdk/resourcemanager/dnsresolver/armdnsresolver/README.md) on how to add the SDK to your project and authenticate.

```go
package armdnsresolver_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver"
)

// x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/preview/2020-04-01-preview/examples/VirtualNetworkLink_Delete.json
func ExampleVirtualNetworkLinksClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdnsresolver.NewVirtualNetworkLinksClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<dns-forwarding-ruleset-name>",
		"<virtual-network-link-name>",
		&armdnsresolver.VirtualNetworkLinksClientBeginDeleteOptions{IfMatch: nil})
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
```
