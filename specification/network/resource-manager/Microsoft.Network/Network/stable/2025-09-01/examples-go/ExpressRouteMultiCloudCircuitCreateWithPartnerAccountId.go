package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v11"
)

// Generated from example definition: 2025-09-01/ExpressRouteMultiCloudCircuitCreateWithPartnerAccountId.json
func ExampleExpressRouteCircuitsClient_BeginCreateOrUpdate_createMultiCloudExpressRouteCircuitWithPartnerAccountId() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExpressRouteCircuitsClient().BeginCreateOrUpdate(ctx, "rg1", "circuitName", armnetwork.ExpressRouteCircuit{
		Properties: &armnetwork.ExpressRouteCircuitPropertiesFormat{
			ServiceProviderProperties: &armnetwork.ExpressRouteCircuitServiceProviderProperties{
				ServiceProviderName: to.Ptr("AWS"),
				PeeringLocation:     to.Ptr("uswest2"),
				BandwidthInMbps:     to.Ptr[int32](200),
			},
			PartnerAccountID: to.Ptr("123456789"),
		},
		SKU: &armnetwork.ExpressRouteCircuitSKU{
			Name:   to.Ptr("MultiCloud_MeteredData"),
			Tier:   to.Ptr(armnetwork.ExpressRouteCircuitSKUTierMultiCloud),
			Family: to.Ptr(armnetwork.ExpressRouteCircuitSKUFamilyMeteredData),
		},
		Location: to.Ptr("eastus2euap"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.ExpressRouteCircuitsClientCreateOrUpdateResponse{
	// 	ExpressRouteCircuit: armnetwork.ExpressRouteCircuit{
	// 		Name: to.Ptr("circuitName"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteCircuits/circuitName"),
	// 		Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
	// 		Type: to.Ptr("Microsoft.Network/expressRouteCircuits"),
	// 		Location: to.Ptr("eastus2euap"),
	// 		Properties: &armnetwork.ExpressRouteCircuitPropertiesFormat{
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateUpdating),
	// 			Peerings: []*armnetwork.ExpressRouteCircuitPeering{
	// 			},
	// 			Authorizations: []*armnetwork.ExpressRouteCircuitAuthorization{
	// 			},
	// 			ServiceProviderProperties: &armnetwork.ExpressRouteCircuitServiceProviderProperties{
	// 				ServiceProviderName: to.Ptr("AWS"),
	// 				PeeringLocation: to.Ptr("uswest2"),
	// 				BandwidthInMbps: to.Ptr[int32](200),
	// 			},
	// 			CircuitProvisioningState: to.Ptr("Enabled"),
	// 			AllowClassicOperations: to.Ptr(false),
	// 			GatewayManagerEtag: to.Ptr(""),
	// 			ServiceKey: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			ServiceProviderProvisioningState: to.Ptr(armnetwork.ServiceProviderProvisioningStateNotProvisioned),
	// 			GlobalReachEnabled: to.Ptr(false),
	// 			EnableDirectPortRateLimit: to.Ptr(false),
	// 			PartnerAccountID: to.Ptr("123456789"),
	// 			ResiliencyLevel: to.Ptr(armnetwork.ResiliencyLevelMaximum),
	// 		},
	// 		SKU: &armnetwork.ExpressRouteCircuitSKU{
	// 			Name: to.Ptr("MultiCloud_MeteredData"),
	// 			Tier: to.Ptr(armnetwork.ExpressRouteCircuitSKUTierMultiCloud),
	// 			Family: to.Ptr(armnetwork.ExpressRouteCircuitSKUFamilyMeteredData),
	// 		},
	// 	},
	// }
}
