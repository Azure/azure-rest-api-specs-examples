package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/LoadBalancerGetReduced.json
func ExampleLoadBalancersClient_Get_getLoadBalancerReduced() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLoadBalancersClient().Get(ctx, "rg-name", "lb-name", &armnetwork.LoadBalancersClientGetOptions{
		DetailLevel: to.Ptr(armnetwork.LoadBalancerDetailLevelReduced)})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.LoadBalancersClientGetResponse{
	// 	LoadBalancer: armnetwork.LoadBalancer{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-name/providers/Microsoft.Network/loadBalancers/lb-name"),
	// 		Name: to.Ptr("lb-name"),
	// 		Type: to.Ptr("Microsoft.Network/loadBalancers"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armnetwork.LoadBalancerPropertiesFormat{
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			FrontendIPConfigurations: []*armnetwork.FrontendIPConfiguration{
	// 				{
	// 					Name: to.Ptr("frontendConfig1"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-name/providers/Microsoft.Network/loadBalancers/lb-name/frontendIPConfigurations/frontendConfig1"),
	// 					Properties: &armnetwork.FrontendIPConfigurationPropertiesFormat{
	// 						PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					},
	// 				},
	// 			},
	// 		},
	// 		SKU: &armnetwork.LoadBalancerSKU{
	// 			Name: to.Ptr(armnetwork.LoadBalancerSKUNameStandard),
	// 			Tier: to.Ptr(armnetwork.LoadBalancerSKUTierRegional),
	// 		},
	// 	},
	// }
}
