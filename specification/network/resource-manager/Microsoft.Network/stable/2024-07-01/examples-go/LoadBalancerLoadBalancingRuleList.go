package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c58fa033619b12c7cfa8a0ec5a9bf03bb18869ab/specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/LoadBalancerLoadBalancingRuleList.json
func ExampleLoadBalancerLoadBalancingRulesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLoadBalancerLoadBalancingRulesClient().NewListPager("testrg", "lb1", nil)
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
		// page.LoadBalancerLoadBalancingRuleListResult = armnetwork.LoadBalancerLoadBalancingRuleListResult{
		// 	Value: []*armnetwork.LoadBalancingRule{
		// 		{
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb1/loadBalancingRules/rule1"),
		// 			Name: to.Ptr("rule1"),
		// 			Type: to.Ptr("Microsoft.Network/loadBalancers/loadBalancingRules"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
		// 			Properties: &armnetwork.LoadBalancingRulePropertiesFormat{
		// 				BackendAddressPool: &armnetwork.SubResource{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb1/backendAddressPools/bepool1"),
		// 				},
		// 				BackendPort: to.Ptr[int32](80),
		// 				EnableFloatingIP: to.Ptr(false),
		// 				EnableTCPReset: to.Ptr(true),
		// 				FrontendIPConfiguration: &armnetwork.SubResource{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb1/frontendIPConfigurations/lbfrontend"),
		// 				},
		// 				FrontendPort: to.Ptr[int32](80),
		// 				IdleTimeoutInMinutes: to.Ptr[int32](15),
		// 				LoadDistribution: to.Ptr(armnetwork.LoadDistributionDefault),
		// 				Probe: &armnetwork.SubResource{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb1/probes/probe1"),
		// 				},
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				Protocol: to.Ptr(armnetwork.TransportProtocolTCP),
		// 			},
		// 	}},
		// }
	}
}
