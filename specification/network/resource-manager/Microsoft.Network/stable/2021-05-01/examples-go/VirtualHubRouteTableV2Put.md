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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VirtualHubRouteTableV2Put.json
func ExampleVirtualHubRouteTableV2SClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armnetwork.NewVirtualHubRouteTableV2SClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<virtual-hub-name>",
		"<route-table-name>",
		armnetwork.VirtualHubRouteTableV2{
			Properties: &armnetwork.VirtualHubRouteTableV2Properties{
				AttachedConnections: []*string{
					to.Ptr("All_Vnets")},
				Routes: []*armnetwork.VirtualHubRouteV2{
					{
						DestinationType: to.Ptr("<destination-type>"),
						Destinations: []*string{
							to.Ptr("20.10.0.0/16"),
							to.Ptr("20.20.0.0/16")},
						NextHopType: to.Ptr("<next-hop-type>"),
						NextHops: []*string{
							to.Ptr("10.0.0.68")},
					},
					{
						DestinationType: to.Ptr("<destination-type>"),
						Destinations: []*string{
							to.Ptr("0.0.0.0/0")},
						NextHopType: to.Ptr("<next-hop-type>"),
						NextHops: []*string{
							to.Ptr("10.0.0.68")},
					}},
			},
		},
		&armnetwork.VirtualHubRouteTableV2SClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetwork%2Farmnetwork%2Fv0.5.0/sdk/resourcemanager/network/armnetwork/README.md) on how to add the SDK to your project and authenticate.
