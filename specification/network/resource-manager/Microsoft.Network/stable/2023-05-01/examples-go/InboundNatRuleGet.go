package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/InboundNatRuleGet.json
func ExampleInboundNatRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewInboundNatRulesClient().Get(ctx, "testrg", "lb1", "natRule1.1", &armnetwork.InboundNatRulesClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.InboundNatRule = armnetwork.InboundNatRule{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb1/inboundNatRules/natRule1.1"),
	// 	Name: to.Ptr("natRule1.1"),
	// 	Properties: &armnetwork.InboundNatRulePropertiesFormat{
	// 		BackendIPConfiguration: &armnetwork.InterfaceIPConfiguration{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Compute/virtualMachineScaleSets/vmss1/virtualMachines/1/networkInterfaces/nic1/ipConfigurations/ip1"),
	// 		},
	// 		BackendPort: to.Ptr[int32](3389),
	// 		EnableFloatingIP: to.Ptr(false),
	// 		EnableTCPReset: to.Ptr(true),
	// 		FrontendIPConfiguration: &armnetwork.SubResource{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb1/frontendIPConfigurations/ip1"),
	// 		},
	// 		FrontendPort: to.Ptr[int32](3390),
	// 		IdleTimeoutInMinutes: to.Ptr[int32](4),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		Protocol: to.Ptr(armnetwork.TransportProtocolTCP),
	// 	},
	// }
}
