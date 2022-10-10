package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/VirtualHubRouteTableV2Put.json
func ExampleVirtualHubRouteTableV2SClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewVirtualHubRouteTableV2SClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "virtualHub1", "virtualHubRouteTable1a", armnetwork.VirtualHubRouteTableV2{
		Properties: &armnetwork.VirtualHubRouteTableV2Properties{
			AttachedConnections: []*string{
				to.Ptr("All_Vnets")},
			Routes: []*armnetwork.VirtualHubRouteV2{
				{
					DestinationType: to.Ptr("CIDR"),
					Destinations: []*string{
						to.Ptr("20.10.0.0/16"),
						to.Ptr("20.20.0.0/16")},
					NextHopType: to.Ptr("IPAddress"),
					NextHops: []*string{
						to.Ptr("10.0.0.68")},
				},
				{
					DestinationType: to.Ptr("CIDR"),
					Destinations: []*string{
						to.Ptr("0.0.0.0/0")},
					NextHopType: to.Ptr("IPAddress"),
					NextHops: []*string{
						to.Ptr("10.0.0.68")},
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
