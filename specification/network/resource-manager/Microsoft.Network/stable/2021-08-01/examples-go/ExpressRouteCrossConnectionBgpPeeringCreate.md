Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetwork%2Farmnetwork%2Fv1.0.0/sdk/resourcemanager/network/armnetwork/README.md) on how to add the SDK to your project and authenticate.

```go
package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/ExpressRouteCrossConnectionBgpPeeringCreate.json
func ExampleExpressRouteCrossConnectionPeeringsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewExpressRouteCrossConnectionPeeringsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"CrossConnection-SiliconValley",
		"<circuitServiceKey>",
		"AzurePrivatePeering",
		armnetwork.ExpressRouteCrossConnectionPeering{
			Properties: &armnetwork.ExpressRouteCrossConnectionPeeringProperties{
				IPv6PeeringConfig: &armnetwork.IPv6ExpressRouteCircuitPeeringConfig{
					PrimaryPeerAddressPrefix:   to.Ptr("3FFE:FFFF:0:CD30::/126"),
					SecondaryPeerAddressPrefix: to.Ptr("3FFE:FFFF:0:CD30::4/126"),
				},
				PeerASN:                    to.Ptr[int64](200),
				PrimaryPeerAddressPrefix:   to.Ptr("192.168.16.252/30"),
				SecondaryPeerAddressPrefix: to.Ptr("192.168.18.252/30"),
				VlanID:                     to.Ptr[int32](200),
			},
		},
		nil)
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
