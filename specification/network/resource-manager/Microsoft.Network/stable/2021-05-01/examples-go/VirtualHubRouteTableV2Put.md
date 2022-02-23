Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetwork%2Farmnetwork%2Fv0.3.1/sdk/resourcemanager/network/armnetwork/README.md) on how to add the SDK to your project and authenticate.

```go
package armnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VirtualHubRouteTableV2Put.json
func ExampleVirtualHubRouteTableV2SClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewVirtualHubRouteTableV2SClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<virtual-hub-name>",
		"<route-table-name>",
		armnetwork.VirtualHubRouteTableV2{
			Properties: &armnetwork.VirtualHubRouteTableV2Properties{
				AttachedConnections: []*string{
					to.StringPtr("All_Vnets")},
				Routes: []*armnetwork.VirtualHubRouteV2{
					{
						DestinationType: to.StringPtr("<destination-type>"),
						Destinations: []*string{
							to.StringPtr("20.10.0.0/16"),
							to.StringPtr("20.20.0.0/16")},
						NextHopType: to.StringPtr("<next-hop-type>"),
						NextHops: []*string{
							to.StringPtr("10.0.0.68")},
					},
					{
						DestinationType: to.StringPtr("<destination-type>"),
						Destinations: []*string{
							to.StringPtr("0.0.0.0/0")},
						NextHopType: to.StringPtr("<next-hop-type>"),
						NextHops: []*string{
							to.StringPtr("10.0.0.68")},
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.VirtualHubRouteTableV2SClientCreateOrUpdateResult)
}
```
