package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NetworkInterfaceEffectiveNSGList.json
func ExampleInterfacesClient_BeginListEffectiveNetworkSecurityGroups() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewInterfacesClient().BeginListEffectiveNetworkSecurityGroups(ctx, "rg1", "nic1", nil)
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
	// res.EffectiveNetworkSecurityGroupListResult = armnetwork.EffectiveNetworkSecurityGroupListResult{
	// 	Value: []*armnetwork.EffectiveNetworkSecurityGroup{
	// 		{
	// 			Association: &armnetwork.EffectiveNetworkSecurityGroupAssociation{
	// 				NetworkInterface: &armnetwork.SubResource{
	// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/nic1"),
	// 				},
	// 				NetworkManager: &armnetwork.SubResource{
	// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/nm1"),
	// 				},
	// 				Subnet: &armnetwork.SubResource{
	// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/rg1-vnet/subnets/default"),
	// 				},
	// 			},
	// 			EffectiveSecurityRules: []*armnetwork.EffectiveNetworkSecurityRule{
	// 				{
	// 					Name: to.Ptr("securityRules/rule1"),
	// 					Access: to.Ptr(armnetwork.SecurityRuleAccessAllow),
	// 					DestinationAddressPrefix: to.Ptr("0.0.0.0/32"),
	// 					DestinationPortRange: to.Ptr("6579-6579"),
	// 					Direction: to.Ptr(armnetwork.SecurityRuleDirectionInbound),
	// 					Priority: to.Ptr[int32](234),
	// 					SourceAddressPrefix: to.Ptr("0.0.0.0/32"),
	// 					SourcePortRange: to.Ptr("456-456"),
	// 					Protocol: to.Ptr(armnetwork.EffectiveSecurityRuleProtocolTCP),
	// 				},
	// 				{
	// 					Name: to.Ptr("securityRules/default-allow-rdp"),
	// 					Access: to.Ptr(armnetwork.SecurityRuleAccessAllow),
	// 					DestinationAddressPrefix: to.Ptr("0.0.0.0/0"),
	// 					DestinationPortRange: to.Ptr("3389-3389"),
	// 					Direction: to.Ptr(armnetwork.SecurityRuleDirectionInbound),
	// 					Priority: to.Ptr[int32](1000),
	// 					SourceAddressPrefix: to.Ptr("1.1.1.1/32"),
	// 					SourcePortRange: to.Ptr("0-65535"),
	// 					Protocol: to.Ptr(armnetwork.EffectiveSecurityRuleProtocolTCP),
	// 				},
	// 				{
	// 					Name: to.Ptr("defaultSecurityRules/AllowInternetOutBound"),
	// 					Access: to.Ptr(armnetwork.SecurityRuleAccessAllow),
	// 					DestinationAddressPrefix: to.Ptr("Internet"),
	// 					DestinationPortRange: to.Ptr("0-65535"),
	// 					Direction: to.Ptr(armnetwork.SecurityRuleDirectionOutbound),
	// 					ExpandedDestinationAddressPrefix: []*string{
	// 						to.Ptr("32.0.0.0/3"),
	// 						to.Ptr("4.0.0.0/6"),
	// 						to.Ptr("2.0.0.0/7"),
	// 						to.Ptr("1.0.0.0/8")},
	// 						Priority: to.Ptr[int32](65001),
	// 						SourceAddressPrefix: to.Ptr("0.0.0.0/0"),
	// 						SourcePortRange: to.Ptr("0-65535"),
	// 						Protocol: to.Ptr(armnetwork.EffectiveSecurityRuleProtocolAll),
	// 				}},
	// 				NetworkSecurityGroup: &armnetwork.SubResource{
	// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/test-nsg"),
	// 				},
	// 		}},
	// 	}
}
