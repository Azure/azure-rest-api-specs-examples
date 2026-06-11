package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/LoadBalancerOutboundRuleList.json
func ExampleLoadBalancerOutboundRulesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLoadBalancerOutboundRulesClient().NewListPager("testrg", "lb1", nil)
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
		// page = armnetwork.LoadBalancerOutboundRulesClientListResponse{
		// 	LoadBalancerOutboundRuleListResult: armnetwork.LoadBalancerOutboundRuleListResult{
		// 		Value: []*armnetwork.OutboundRule{
		// 			{
		// 				Name: to.Ptr("rule1"),
		// 				Type: to.Ptr("Microsoft.Network/loadBalancers/outboundRules"),
		// 				Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb1/outboundRules/rule1"),
		// 				Properties: &armnetwork.OutboundRulePropertiesFormat{
		// 					AllocatedOutboundPorts: to.Ptr[int32](64),
		// 					BackendAddressPool: &armnetwork.SubResource{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb1/backendAddressPools/bepool1"),
		// 					},
		// 					EnableTCPReset: to.Ptr(true),
		// 					FrontendIPConfigurations: []*armnetwork.SubResource{
		// 						{
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb1/frontendIPConfigurations/lbfrontend"),
		// 						},
		// 					},
		// 					IdleTimeoutInMinutes: to.Ptr[int32](15),
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					Protocol: to.Ptr(armnetwork.LoadBalancerOutboundRuleProtocolTCP),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
