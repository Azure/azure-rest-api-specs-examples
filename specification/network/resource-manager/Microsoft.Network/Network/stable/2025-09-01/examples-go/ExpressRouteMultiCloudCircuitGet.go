package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v11"
)

// Generated from example definition: 2025-09-01/ExpressRouteMultiCloudCircuitGet.json
func ExampleExpressRouteCircuitsClient_Get_getMultiCloudExpressRouteCircuit() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExpressRouteCircuitsClient().Get(ctx, "rg1", "circuitName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.ExpressRouteCircuitsClientGetResponse{
	// 	ExpressRouteCircuit: armnetwork.ExpressRouteCircuit{
	// 		Name: to.Ptr("circuitName"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteCircuits/circuitName"),
	// 		Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
	// 		Type: to.Ptr("Microsoft.Network/expressRouteCircuits"),
	// 		Location: to.Ptr("eastus2euap"),
	// 		Properties: &armnetwork.ExpressRouteCircuitPropertiesFormat{
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			Authorizations: []*armnetwork.ExpressRouteCircuitAuthorization{
	// 			},
	// 			ServiceProviderProperties: &armnetwork.ExpressRouteCircuitServiceProviderProperties{
	// 				ServiceProviderName: to.Ptr("AWS"),
	// 				PeeringLocation: to.Ptr("uswest2"),
	// 				BandwidthInMbps: to.Ptr[int32](200),
	// 			},
	// 			PartnerAccountID: to.Ptr("123-456-789"),
	// 			ActivationKey: to.Ptr("ew0KICAic2hhcmVkQ29ubmVjdGlvblV1aWQiOiAiM2YxYTZjOWUtOGIzZi00YjllLTljM2EtMmU2ZDdmNGExYzU5IiwNCiAgImNvbm5lY3Rpb25TaXplTWJwcyI6IDIwMCwNCiAgImRlc3RpbmF0aW9uQWNjb3VudCI6ICI3ZDc0N2VlZC1iNDRjLTQyNTctOGQ0My1kZjllYmQ5NDU0NmIiDQp9"),
	// 			ResiliencyLevel: to.Ptr(armnetwork.ResiliencyLevelMaximum),
	// 			CircuitProvisioningState: to.Ptr("Enabled"),
	// 			AllowClassicOperations: to.Ptr(false),
	// 			ServiceKey: to.Ptr("a1410692-0000-4ceb-b94a-b90b94d398d1"),
	// 			ServiceProviderProvisioningState: to.Ptr(armnetwork.ServiceProviderProvisioningStateDeProvisioned),
	// 		},
	// 		SKU: &armnetwork.ExpressRouteCircuitSKU{
	// 			Name: to.Ptr("MultiCloud_MeteredData"),
	// 			Tier: to.Ptr(armnetwork.ExpressRouteCircuitSKUTierMultiCloud),
	// 			Family: to.Ptr(armnetwork.ExpressRouteCircuitSKUFamilyMeteredData),
	// 		},
	// 	},
	// }
}
