package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/NetworkSecurityGroupList.json
func ExampleSecurityGroupsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSecurityGroupsClient().NewListPager("rg1", nil)
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
		// page = armnetwork.SecurityGroupsClientListResponse{
		// 	SecurityGroupListResult: armnetwork.SecurityGroupListResult{
		// 		Value: []*armnetwork.SecurityGroup{
		// 			{
		// 				Name: to.Ptr("nsg1"),
		// 				Type: to.Ptr("Microsoft.Network/networkSecurityGroups"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/nsg1"),
		// 				Location: to.Ptr("westus"),
		// 				Properties: &armnetwork.SecurityGroupPropertiesFormat{
		// 					DefaultSecurityRules: []*armnetwork.SecurityRule{
		// 						{
		// 							Name: to.Ptr("AllowVnetInBound"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/nsg1/defaultSecurityRules/AllowVnetInBound"),
		// 							Properties: &armnetwork.SecurityRulePropertiesFormat{
		// 								Description: to.Ptr("Allow inbound traffic from all VMs in VNET"),
		// 								Access: to.Ptr(armnetwork.SecurityRuleAccessAllow),
		// 								DestinationAddressPrefix: to.Ptr("VirtualNetwork"),
		// 								DestinationPortRange: to.Ptr("*"),
		// 								Direction: to.Ptr(armnetwork.SecurityRuleDirectionInbound),
		// 								Priority: to.Ptr[int32](65000),
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								SourceAddressPrefix: to.Ptr("VirtualNetwork"),
		// 								SourcePortRange: to.Ptr("*"),
		// 								Protocol: to.Ptr(armnetwork.SecurityRuleProtocolAsterisk),
		// 							},
		// 						},
		// 						{
		// 							Name: to.Ptr("AllowAzureLoadBalancerInBound"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/nsg1/defaultSecurityRules/AllowAzureLoadBalancerInBound"),
		// 							Properties: &armnetwork.SecurityRulePropertiesFormat{
		// 								Description: to.Ptr("Allow inbound traffic from azure load balancer"),
		// 								Access: to.Ptr(armnetwork.SecurityRuleAccessAllow),
		// 								DestinationAddressPrefix: to.Ptr("*"),
		// 								DestinationPortRange: to.Ptr("*"),
		// 								Direction: to.Ptr(armnetwork.SecurityRuleDirectionInbound),
		// 								Priority: to.Ptr[int32](65001),
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								SourceAddressPrefix: to.Ptr("AzureLoadBalancer"),
		// 								SourcePortRange: to.Ptr("*"),
		// 								Protocol: to.Ptr(armnetwork.SecurityRuleProtocolAsterisk),
		// 							},
		// 						},
		// 						{
		// 							Name: to.Ptr("DenyAllInBound"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/nsg1/defaultSecurityRules/DenyAllInBound"),
		// 							Properties: &armnetwork.SecurityRulePropertiesFormat{
		// 								Description: to.Ptr("Deny all inbound traffic"),
		// 								Access: to.Ptr(armnetwork.SecurityRuleAccessDeny),
		// 								DestinationAddressPrefix: to.Ptr("*"),
		// 								DestinationPortRange: to.Ptr("*"),
		// 								Direction: to.Ptr(armnetwork.SecurityRuleDirectionInbound),
		// 								Priority: to.Ptr[int32](65500),
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								SourceAddressPrefix: to.Ptr("*"),
		// 								SourcePortRange: to.Ptr("*"),
		// 								Protocol: to.Ptr(armnetwork.SecurityRuleProtocolAsterisk),
		// 							},
		// 						},
		// 						{
		// 							Name: to.Ptr("AllowVnetOutBound"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/nsg1/defaultSecurityRules/AllowVnetOutBound"),
		// 							Properties: &armnetwork.SecurityRulePropertiesFormat{
		// 								Description: to.Ptr("Allow outbound traffic from all VMs to all VMs in VNET"),
		// 								Access: to.Ptr(armnetwork.SecurityRuleAccessAllow),
		// 								DestinationAddressPrefix: to.Ptr("VirtualNetwork"),
		// 								DestinationPortRange: to.Ptr("*"),
		// 								Direction: to.Ptr(armnetwork.SecurityRuleDirectionOutbound),
		// 								Priority: to.Ptr[int32](65000),
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								SourceAddressPrefix: to.Ptr("VirtualNetwork"),
		// 								SourcePortRange: to.Ptr("*"),
		// 								Protocol: to.Ptr(armnetwork.SecurityRuleProtocolAsterisk),
		// 							},
		// 						},
		// 						{
		// 							Name: to.Ptr("AllowInternetOutBound"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/nsg1/defaultSecurityRules/AllowInternetOutBound"),
		// 							Properties: &armnetwork.SecurityRulePropertiesFormat{
		// 								Description: to.Ptr("Allow outbound traffic from all VMs to Internet"),
		// 								Access: to.Ptr(armnetwork.SecurityRuleAccessAllow),
		// 								DestinationAddressPrefix: to.Ptr("Internet"),
		// 								DestinationPortRange: to.Ptr("*"),
		// 								Direction: to.Ptr(armnetwork.SecurityRuleDirectionOutbound),
		// 								Priority: to.Ptr[int32](65001),
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								SourceAddressPrefix: to.Ptr("*"),
		// 								SourcePortRange: to.Ptr("*"),
		// 								Protocol: to.Ptr(armnetwork.SecurityRuleProtocolAsterisk),
		// 							},
		// 						},
		// 						{
		// 							Name: to.Ptr("DenyAllOutBound"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/nsg1/defaultSecurityRules/DenyAllOutBound"),
		// 							Properties: &armnetwork.SecurityRulePropertiesFormat{
		// 								Description: to.Ptr("Deny all outbound traffic"),
		// 								Access: to.Ptr(armnetwork.SecurityRuleAccessDeny),
		// 								DestinationAddressPrefix: to.Ptr("*"),
		// 								DestinationPortRange: to.Ptr("*"),
		// 								Direction: to.Ptr(armnetwork.SecurityRuleDirectionOutbound),
		// 								Priority: to.Ptr[int32](65500),
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								SourceAddressPrefix: to.Ptr("*"),
		// 								SourcePortRange: to.Ptr("*"),
		// 								Protocol: to.Ptr(armnetwork.SecurityRuleProtocolAsterisk),
		// 							},
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					SecurityRules: []*armnetwork.SecurityRule{
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("nsg3"),
		// 				Type: to.Ptr("Microsoft.Network/networkSecurityGroups"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/nsg3"),
		// 				Location: to.Ptr("westus"),
		// 				Properties: &armnetwork.SecurityGroupPropertiesFormat{
		// 					DefaultSecurityRules: []*armnetwork.SecurityRule{
		// 						{
		// 							Name: to.Ptr("AllowVnetInBound"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/nsg3/defaultSecurityRules/AllowVnetInBound"),
		// 							Properties: &armnetwork.SecurityRulePropertiesFormat{
		// 								Description: to.Ptr("Allow inbound traffic from all VMs in VNET"),
		// 								Access: to.Ptr(armnetwork.SecurityRuleAccessAllow),
		// 								DestinationAddressPrefix: to.Ptr("VirtualNetwork"),
		// 								DestinationPortRange: to.Ptr("*"),
		// 								Direction: to.Ptr(armnetwork.SecurityRuleDirectionInbound),
		// 								Priority: to.Ptr[int32](65000),
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								SourceAddressPrefix: to.Ptr("VirtualNetwork"),
		// 								SourcePortRange: to.Ptr("*"),
		// 								Protocol: to.Ptr(armnetwork.SecurityRuleProtocolAsterisk),
		// 							},
		// 						},
		// 						{
		// 							Name: to.Ptr("AllowAzureLoadBalancerInBound"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/nsg3/defaultSecurityRules/AllowAzureLoadBalancerInBound"),
		// 							Properties: &armnetwork.SecurityRulePropertiesFormat{
		// 								Description: to.Ptr("Allow inbound traffic from azure load balancer"),
		// 								Access: to.Ptr(armnetwork.SecurityRuleAccessAllow),
		// 								DestinationAddressPrefix: to.Ptr("*"),
		// 								DestinationPortRange: to.Ptr("*"),
		// 								Direction: to.Ptr(armnetwork.SecurityRuleDirectionInbound),
		// 								Priority: to.Ptr[int32](65001),
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								SourceAddressPrefix: to.Ptr("AzureLoadBalancer"),
		// 								SourcePortRange: to.Ptr("*"),
		// 								Protocol: to.Ptr(armnetwork.SecurityRuleProtocolAsterisk),
		// 							},
		// 						},
		// 						{
		// 							Name: to.Ptr("DenyAllInBound"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/nsg3/defaultSecurityRules/DenyAllInBound"),
		// 							Properties: &armnetwork.SecurityRulePropertiesFormat{
		// 								Description: to.Ptr("Deny all inbound traffic"),
		// 								Access: to.Ptr(armnetwork.SecurityRuleAccessDeny),
		// 								DestinationAddressPrefix: to.Ptr("*"),
		// 								DestinationPortRange: to.Ptr("*"),
		// 								Direction: to.Ptr(armnetwork.SecurityRuleDirectionInbound),
		// 								Priority: to.Ptr[int32](65500),
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								SourceAddressPrefix: to.Ptr("*"),
		// 								SourcePortRange: to.Ptr("*"),
		// 								Protocol: to.Ptr(armnetwork.SecurityRuleProtocolAsterisk),
		// 							},
		// 						},
		// 						{
		// 							Name: to.Ptr("AllowVnetOutBound"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/nsg3/defaultSecurityRules/AllowVnetOutBound"),
		// 							Properties: &armnetwork.SecurityRulePropertiesFormat{
		// 								Description: to.Ptr("Allow outbound traffic from all VMs to all VMs in VNET"),
		// 								Access: to.Ptr(armnetwork.SecurityRuleAccessAllow),
		// 								DestinationAddressPrefix: to.Ptr("VirtualNetwork"),
		// 								DestinationPortRange: to.Ptr("*"),
		// 								Direction: to.Ptr(armnetwork.SecurityRuleDirectionOutbound),
		// 								Priority: to.Ptr[int32](65000),
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								SourceAddressPrefix: to.Ptr("VirtualNetwork"),
		// 								SourcePortRange: to.Ptr("*"),
		// 								Protocol: to.Ptr(armnetwork.SecurityRuleProtocolAsterisk),
		// 							},
		// 						},
		// 						{
		// 							Name: to.Ptr("AllowInternetOutBound"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/nsg3/defaultSecurityRules/AllowInternetOutBound"),
		// 							Properties: &armnetwork.SecurityRulePropertiesFormat{
		// 								Description: to.Ptr("Allow outbound traffic from all VMs to Internet"),
		// 								Access: to.Ptr(armnetwork.SecurityRuleAccessAllow),
		// 								DestinationAddressPrefix: to.Ptr("Internet"),
		// 								DestinationPortRange: to.Ptr("*"),
		// 								Direction: to.Ptr(armnetwork.SecurityRuleDirectionOutbound),
		// 								Priority: to.Ptr[int32](65001),
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								SourceAddressPrefix: to.Ptr("*"),
		// 								SourcePortRange: to.Ptr("*"),
		// 								Protocol: to.Ptr(armnetwork.SecurityRuleProtocolAsterisk),
		// 							},
		// 						},
		// 						{
		// 							Name: to.Ptr("DenyAllOutBound"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/nsg3/defaultSecurityRules/DenyAllOutBound"),
		// 							Properties: &armnetwork.SecurityRulePropertiesFormat{
		// 								Description: to.Ptr("Deny all outbound traffic"),
		// 								Access: to.Ptr(armnetwork.SecurityRuleAccessDeny),
		// 								DestinationAddressPrefix: to.Ptr("*"),
		// 								DestinationPortRange: to.Ptr("*"),
		// 								Direction: to.Ptr(armnetwork.SecurityRuleDirectionOutbound),
		// 								Priority: to.Ptr[int32](65500),
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								SourceAddressPrefix: to.Ptr("*"),
		// 								SourcePortRange: to.Ptr("*"),
		// 								Protocol: to.Ptr(armnetwork.SecurityRuleProtocolAsterisk),
		// 							},
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					SecurityRules: []*armnetwork.SecurityRule{
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
