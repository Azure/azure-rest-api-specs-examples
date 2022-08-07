package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/ExpressRoutePortUpdateLink.json
func ExampleExpressRoutePortsClient_BeginCreateOrUpdate_expressRoutePortUpdateLink() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewExpressRoutePortsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "portName", armnetwork.ExpressRoutePort{
		Location: to.Ptr("westus"),
		Properties: &armnetwork.ExpressRoutePortPropertiesFormat{
			BandwidthInGbps: to.Ptr[int32](100),
			Encapsulation:   to.Ptr(armnetwork.ExpressRoutePortsEncapsulationQinQ),
			Links: []*armnetwork.ExpressRouteLink{
				{
					Name: to.Ptr("link1"),
					Properties: &armnetwork.ExpressRouteLinkPropertiesFormat{
						AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateEnabled),
					},
				}},
			PeeringLocation: to.Ptr("peeringLocationName"),
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
