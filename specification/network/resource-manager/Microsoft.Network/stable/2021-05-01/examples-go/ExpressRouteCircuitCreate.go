package armnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/ExpressRouteCircuitCreate.json
func ExampleExpressRouteCircuitsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armnetwork.NewExpressRouteCircuitsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<circuit-name>",
		armnetwork.ExpressRouteCircuit{
			Location: to.Ptr("<location>"),
			Properties: &armnetwork.ExpressRouteCircuitPropertiesFormat{
				AllowClassicOperations: to.Ptr(false),
				Authorizations:         []*armnetwork.ExpressRouteCircuitAuthorization{},
				Peerings:               []*armnetwork.ExpressRouteCircuitPeering{},
				ServiceProviderProperties: &armnetwork.ExpressRouteCircuitServiceProviderProperties{
					BandwidthInMbps:     to.Ptr[int32](200),
					PeeringLocation:     to.Ptr("<peering-location>"),
					ServiceProviderName: to.Ptr("<service-provider-name>"),
				},
			},
			SKU: &armnetwork.ExpressRouteCircuitSKU{
				Name:   to.Ptr("<name>"),
				Family: to.Ptr(armnetwork.ExpressRouteCircuitSKUFamilyMeteredData),
				Tier:   to.Ptr(armnetwork.ExpressRouteCircuitSKUTierStandard),
			},
		},
		&armnetwork.ExpressRouteCircuitsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
