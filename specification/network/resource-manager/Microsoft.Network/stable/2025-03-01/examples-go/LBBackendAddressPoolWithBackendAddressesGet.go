package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/LBBackendAddressPoolWithBackendAddressesGet.json
func ExampleLoadBalancerBackendAddressPoolsClient_Get_loadBalancerWithBackendAddressPoolWithBackendAddresses() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLoadBalancerBackendAddressPoolsClient().Get(ctx, "testrg", "lb", "backend", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BackendAddressPool = armnetwork.BackendAddressPool{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb/backendAddressPools/backend"),
	// 	Name: to.Ptr("backend"),
	// 	Type: to.Ptr("Microsoft.Network/loadBalancers/backendAddressPools"),
	// 	Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
	// 	Properties: &armnetwork.BackendAddressPoolPropertiesFormat{
	// 		LoadBalancerBackendAddresses: []*armnetwork.LoadBalancerBackendAddress{
	// 			{
	// 				Name: to.Ptr("address1"),
	// 				Properties: &armnetwork.LoadBalancerBackendAddressPropertiesFormat{
	// 					IPAddress: to.Ptr("10.0.0.4"),
	// 					VirtualNetwork: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb"),
	// 					},
	// 				},
	// 			},
	// 			{
	// 				Name: to.Ptr("address2"),
	// 				Properties: &armnetwork.LoadBalancerBackendAddressPropertiesFormat{
	// 					IPAddress: to.Ptr("10.0.0.5"),
	// 					VirtualNetwork: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb"),
	// 					},
	// 				},
	// 		}},
	// 		LoadBalancingRules: []*armnetwork.SubResource{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb/loadBalancingRules/rulelb"),
	// 		}},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 	},
	// }
}
