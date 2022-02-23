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

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/HubRouteTablePut.json
func ExampleHubRouteTablesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewHubRouteTablesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<virtual-hub-name>",
		"<route-table-name>",
		armnetwork.HubRouteTable{
			Properties: &armnetwork.HubRouteTableProperties{
				Labels: []*string{
					to.StringPtr("label1"),
					to.StringPtr("label2")},
				Routes: []*armnetwork.HubRoute{
					{
						Name:            to.StringPtr("<name>"),
						DestinationType: to.StringPtr("<destination-type>"),
						Destinations: []*string{
							to.StringPtr("10.0.0.0/8"),
							to.StringPtr("20.0.0.0/8"),
							to.StringPtr("30.0.0.0/8")},
						NextHop:     to.StringPtr("<next-hop>"),
						NextHopType: to.StringPtr("<next-hop-type>"),
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
	log.Printf("Response result: %#v\n", res.HubRouteTablesClientCreateOrUpdateResult)
}
```
