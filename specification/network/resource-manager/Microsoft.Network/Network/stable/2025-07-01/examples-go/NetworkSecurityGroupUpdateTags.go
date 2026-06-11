package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/NetworkSecurityGroupUpdateTags.json
func ExampleSecurityGroupsClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSecurityGroupsClient().UpdateTags(ctx, "rg1", "testnsg", armnetwork.TagsObject{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.SecurityGroupsClientUpdateTagsResponse{
	// 	SecurityGroup: armnetwork.SecurityGroup{
	// 		Name: to.Ptr("testnsg"),
	// 		Type: to.Ptr("Microsoft.Network/networkSecurityGroups"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/testnsg"),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armnetwork.SecurityGroupPropertiesFormat{
	// 			DefaultSecurityRules: []*armnetwork.SecurityRule{
	// 				{
	// 					Name: to.Ptr("AllowVnetInBound"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/testnsg/defaultSecurityRules/AllowVnetInBound"),
	// 					Properties: &armnetwork.SecurityRulePropertiesFormat{
	// 						Description: to.Ptr("Allow inbound traffic from all VMs in VNET"),
	// 						Access: to.Ptr(armnetwork.SecurityRuleAccessAllow),
	// 						DestinationAddressPrefix: to.Ptr("VirtualNetwork"),
	// 						DestinationPortRange: to.Ptr("*"),
	// 						Direction: to.Ptr(armnetwork.SecurityRuleDirectionInbound),
	// 						Priority: to.Ptr[int32](65000),
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 						SourceAddressPrefix: to.Ptr("VirtualNetwork"),
	// 						SourcePortRange: to.Ptr("*"),
	// 						Protocol: to.Ptr(armnetwork.SecurityRuleProtocolAsterisk),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("AllowAzureLoadBalancerInBound"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/testnsg/defaultSecurityRules/AllowAzureLoadBalancerInBound"),
	// 					Properties: &armnetwork.SecurityRulePropertiesFormat{
	// 						Description: to.Ptr("Allow inbound traffic from azure load balancer"),
	// 						Access: to.Ptr(armnetwork.SecurityRuleAccessAllow),
	// 						DestinationAddressPrefix: to.Ptr("*"),
	// 						DestinationPortRange: to.Ptr("*"),
	// 						Direction: to.Ptr(armnetwork.SecurityRuleDirectionInbound),
	// 						Priority: to.Ptr[int32](65001),
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 						SourceAddressPrefix: to.Ptr("AzureLoadBalancer"),
	// 						SourcePortRange: to.Ptr("*"),
	// 						Protocol: to.Ptr(armnetwork.SecurityRuleProtocolAsterisk),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("DenyAllInBound"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/testnsg/defaultSecurityRules/DenyAllInBound"),
	// 					Properties: &armnetwork.SecurityRulePropertiesFormat{
	// 						Description: to.Ptr("Deny all inbound traffic"),
	// 						Access: to.Ptr(armnetwork.SecurityRuleAccessDeny),
	// 						DestinationAddressPrefix: to.Ptr("*"),
	// 						DestinationPortRange: to.Ptr("*"),
	// 						Direction: to.Ptr(armnetwork.SecurityRuleDirectionInbound),
	// 						Priority: to.Ptr[int32](65500),
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 						SourceAddressPrefix: to.Ptr("*"),
	// 						SourcePortRange: to.Ptr("*"),
	// 						Protocol: to.Ptr(armnetwork.SecurityRuleProtocolAsterisk),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("AllowVnetOutBound"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/testnsg/defaultSecurityRules/AllowVnetOutBound"),
	// 					Properties: &armnetwork.SecurityRulePropertiesFormat{
	// 						Description: to.Ptr("Allow outbound traffic from all VMs to all VMs in VNET"),
	// 						Access: to.Ptr(armnetwork.SecurityRuleAccessAllow),
	// 						DestinationAddressPrefix: to.Ptr("VirtualNetwork"),
	// 						DestinationPortRange: to.Ptr("*"),
	// 						Direction: to.Ptr(armnetwork.SecurityRuleDirectionOutbound),
	// 						Priority: to.Ptr[int32](65000),
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 						SourceAddressPrefix: to.Ptr("VirtualNetwork"),
	// 						SourcePortRange: to.Ptr("*"),
	// 						Protocol: to.Ptr(armnetwork.SecurityRuleProtocolAsterisk),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("AllowInternetOutBound"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/testnsg/defaultSecurityRules/AllowInternetOutBound"),
	// 					Properties: &armnetwork.SecurityRulePropertiesFormat{
	// 						Description: to.Ptr("Allow outbound traffic from all VMs to Internet"),
	// 						Access: to.Ptr(armnetwork.SecurityRuleAccessAllow),
	// 						DestinationAddressPrefix: to.Ptr("Internet"),
	// 						DestinationPortRange: to.Ptr("*"),
	// 						Direction: to.Ptr(armnetwork.SecurityRuleDirectionOutbound),
	// 						Priority: to.Ptr[int32](65001),
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 						SourceAddressPrefix: to.Ptr("*"),
	// 						SourcePortRange: to.Ptr("*"),
	// 						Protocol: to.Ptr(armnetwork.SecurityRuleProtocolAsterisk),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("DenyAllOutBound"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/testnsg/defaultSecurityRules/DenyAllOutBound"),
	// 					Properties: &armnetwork.SecurityRulePropertiesFormat{
	// 						Description: to.Ptr("Deny all outbound traffic"),
	// 						Access: to.Ptr(armnetwork.SecurityRuleAccessDeny),
	// 						DestinationAddressPrefix: to.Ptr("*"),
	// 						DestinationPortRange: to.Ptr("*"),
	// 						Direction: to.Ptr(armnetwork.SecurityRuleDirectionOutbound),
	// 						Priority: to.Ptr[int32](65500),
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 						SourceAddressPrefix: to.Ptr("*"),
	// 						SourcePortRange: to.Ptr("*"),
	// 						Protocol: to.Ptr(armnetwork.SecurityRuleProtocolAsterisk),
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			SecurityRules: []*armnetwork.SecurityRule{
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"tag1": to.Ptr("value1"),
	// 			"tag2": to.Ptr("value2"),
	// 		},
	// 	},
	// }
}
