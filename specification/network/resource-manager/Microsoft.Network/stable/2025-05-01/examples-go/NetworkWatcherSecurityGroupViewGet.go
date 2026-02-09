package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NetworkWatcherSecurityGroupViewGet.json
func ExampleWatchersClient_BeginGetVMSecurityRules() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWatchersClient().BeginGetVMSecurityRules(ctx, "rg1", "nw1", armnetwork.SecurityGroupViewParameters{
		TargetResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Compute/virtualMachines/vm1"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SecurityGroupViewResult = armnetwork.SecurityGroupViewResult{
	// 	NetworkInterfaces: []*armnetwork.SecurityGroupNetworkInterface{
	// 		{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/networkInterfaces/nic1"),
	// 			SecurityRuleAssociations: &armnetwork.SecurityRuleAssociations{
	// 				DefaultSecurityRules: []*armnetwork.SecurityRule{
	// 					{
	// 						ID: to.Ptr("/subscriptions//resourceGroups//providers/Microsoft.Network/networkSecurityGroups//defaultSecurityRules/"),
	// 						Name: to.Ptr("AllowVnetInBound"),
	// 						Properties: &armnetwork.SecurityRulePropertiesFormat{
	// 							Description: to.Ptr("Allow inbound traffic from all VMs in VNET"),
	// 							Access: to.Ptr(armnetwork.SecurityRuleAccessAllow),
	// 							DestinationAddressPrefix: to.Ptr("VirtualNetwork"),
	// 							DestinationPortRange: to.Ptr("*"),
	// 							Direction: to.Ptr(armnetwork.SecurityRuleDirectionInbound),
	// 							Priority: to.Ptr[int32](65000),
	// 							ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 							SourceAddressPrefix: to.Ptr("VirtualNetwork"),
	// 							SourcePortRange: to.Ptr("*"),
	// 							Protocol: to.Ptr(armnetwork.SecurityRuleProtocolAsterisk),
	// 						},
	// 				}},
	// 				EffectiveSecurityRules: []*armnetwork.EffectiveNetworkSecurityRule{
	// 					{
	// 						Name: to.Ptr("DefaultOutboundDenyAll"),
	// 						Access: to.Ptr(armnetwork.SecurityRuleAccessDeny),
	// 						DestinationAddressPrefix: to.Ptr("*"),
	// 						DestinationPortRange: to.Ptr("0-65535"),
	// 						Direction: to.Ptr(armnetwork.SecurityRuleDirectionOutbound),
	// 						Priority: to.Ptr[int32](65500),
	// 						SourceAddressPrefix: to.Ptr("*"),
	// 						SourcePortRange: to.Ptr("0-65535"),
	// 						Protocol: to.Ptr(armnetwork.EffectiveSecurityRuleProtocolAll),
	// 				}},
	// 				SubnetAssociation: &armnetwork.SubnetAssociation{
	// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
	// 					SecurityRules: []*armnetwork.SecurityRule{
	// 						{
	// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/networkSecurityGroups/AppNSG/securityRules/fe_rule"),
	// 							Name: to.Ptr("fe_rule"),
	// 							Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
	// 							Properties: &armnetwork.SecurityRulePropertiesFormat{
	// 								Description: to.Ptr("Allow Frontend"),
	// 								Access: to.Ptr(armnetwork.SecurityRuleAccessAllow),
	// 								DestinationAddressPrefix: to.Ptr("*"),
	// 								DestinationPortRange: to.Ptr("*"),
	// 								Direction: to.Ptr(armnetwork.SecurityRuleDirectionInbound),
	// 								Priority: to.Ptr[int32](100),
	// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 								SourceAddressPrefix: to.Ptr("10.1.0.0/24"),
	// 								SourcePortRange: to.Ptr("*"),
	// 								Protocol: to.Ptr(armnetwork.SecurityRuleProtocolTCP),
	// 							},
	// 					}},
	// 				},
	// 			},
	// 	}},
	// }
}
