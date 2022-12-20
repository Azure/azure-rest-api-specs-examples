package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/HubRouteTablePut.json
func ExampleHubRouteTablesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewHubRouteTablesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "virtualHub1", "hubRouteTable1", armnetwork.HubRouteTable{
		Properties: &armnetwork.HubRouteTableProperties{
			Labels: []*string{
				to.Ptr("label1"),
				to.Ptr("label2")},
			Routes: []*armnetwork.HubRoute{
				{
					Name:            to.Ptr("route1"),
					DestinationType: to.Ptr("CIDR"),
					Destinations: []*string{
						to.Ptr("10.0.0.0/8"),
						to.Ptr("20.0.0.0/8"),
						to.Ptr("30.0.0.0/8")},
					NextHop:     to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azureFirewall1"),
					NextHopType: to.Ptr("ResourceId"),
				}},
		},
	}, nil)
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
