package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/InboundNatRuleList.json
func ExampleInboundNatRulesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewInboundNatRulesClient().NewListPager("testrg", "lb1", nil)
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
		// page = armnetwork.InboundNatRulesClientListResponse{
		// 	InboundNatRuleListResult: armnetwork.InboundNatRuleListResult{
		// 		Value: []*armnetwork.InboundNatRule{
		// 			{
		// 				Name: to.Ptr("natRule1.1"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb1/inboundNatRules/natRule1.1"),
		// 				Properties: &armnetwork.InboundNatRulePropertiesFormat{
		// 					BackendIPConfiguration: &armnetwork.InterfaceIPConfiguration{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Compute/virtualMachineScaleSets/vmss1/virtualMachines/1/networkInterfaces/nic1/ipConfigurations/ip1"),
		// 					},
		// 					BackendPort: to.Ptr[int32](3389),
		// 					EnableFloatingIP: to.Ptr(false),
		// 					EnableTCPReset: to.Ptr(true),
		// 					FrontendIPConfiguration: &armnetwork.SubResource{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb1/frontendIPConfigurations/ip1"),
		// 					},
		// 					FrontendPort: to.Ptr[int32](3390),
		// 					IdleTimeoutInMinutes: to.Ptr[int32](4),
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					Protocol: to.Ptr(armnetwork.TransportProtocolTCP),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("natRule1.3"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb1/inboundNatRules/natRule1.3"),
		// 				Properties: &armnetwork.InboundNatRulePropertiesFormat{
		// 					BackendIPConfiguration: &armnetwork.InterfaceIPConfiguration{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Compute/virtualMachineScaleSets/vmss1/virtualMachines/3/networkInterfaces/nic1/ipConfigurations/ip1"),
		// 					},
		// 					BackendPort: to.Ptr[int32](3389),
		// 					EnableFloatingIP: to.Ptr(false),
		// 					EnableTCPReset: to.Ptr(true),
		// 					FrontendIPConfiguration: &armnetwork.SubResource{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb1/frontendIPConfigurations/ip1"),
		// 					},
		// 					FrontendPort: to.Ptr[int32](3392),
		// 					IdleTimeoutInMinutes: to.Ptr[int32](4),
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					Protocol: to.Ptr(armnetwork.TransportProtocolTCP),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
