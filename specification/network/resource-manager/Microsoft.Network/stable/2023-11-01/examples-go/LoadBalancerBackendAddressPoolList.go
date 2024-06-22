package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f4c6c8697c59f966db0d1e36b62df3af3bca9065/specification/network/resource-manager/Microsoft.Network/stable/2023-11-01/examples/LoadBalancerBackendAddressPoolList.json
func ExampleLoadBalancerBackendAddressPoolsClient_NewListPager_loadBalancerBackendAddressPoolList() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLoadBalancerBackendAddressPoolsClient().NewListPager("testrg", "lb", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.LoadBalancerBackendAddressPoolListResult = armnetwork.LoadBalancerBackendAddressPoolListResult{
		// 	Value: []*armnetwork.BackendAddressPool{
		// 		{
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb/backendAddressPools/backend"),
		// 			Name: to.Ptr("backend"),
		// 			Type: to.Ptr("Microsoft.Network/loadBalancers/backendAddressPools"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
		// 			Properties: &armnetwork.BackendAddressPoolPropertiesFormat{
		// 				BackendIPConfigurations: []*armnetwork.InterfaceIPConfiguration{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/networkInterfaces/nic/ipConfigurations/default-ip-config"),
		// 				}},
		// 				LoadBalancingRules: []*armnetwork.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb/loadBalancingRules/rulelb"),
		// 				}},
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
