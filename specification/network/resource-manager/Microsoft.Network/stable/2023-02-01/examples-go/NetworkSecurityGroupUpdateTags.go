package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9975d3476c05bcc6bd9535ad3dfb564e6a168fa5/specification/network/resource-manager/Microsoft.Network/stable/2023-02-01/examples/NetworkSecurityGroupUpdateTags.json
func ExampleSecurityGroupsClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.SecurityGroup = armnetwork.SecurityGroup{
	// 	Name: to.Ptr("testnsg"),
	// 	Type: to.Ptr("Microsoft.Network/networkSecurityGroups"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/testnsg"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armnetwork.SecurityGroupPropertiesFormat{
	// 		DefaultSecurityRules: []*armnetwork.SecurityRule{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/testnsg/defaultSecurityRules/AllowVnetInBound"),
	// 				Name: to.Ptr("AllowVnetInBound"),
	// 				Properties: &armnetwork.SecurityRulePropertiesFormat{
	// 					Description: to.Ptr("Allow inbound traffic from all VMs in VNET"),
	// 					Access: to.Ptr(armnetwork.SecurityRuleAccessAllow),
	// 					DestinationAddressPrefix: to.Ptr("VirtualNetwork"),
	// 					DestinationPortRange: to.Ptr("*"),
	// 					Direction: to.Ptr(armnetwork.SecurityRuleDirectionInbound),
	// 					Priority: to.Ptr[int32](65000),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					SourceAddressPrefix: to.Ptr("VirtualNetwork"),
	// 					SourcePortRange: to.Ptr("*"),
	// 					Protocol: to.Ptr(armnetwork.SecurityRuleProtocolAsterisk),
	// 				},
	// 			},
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/testnsg/defaultSecurityRules/AllowAzureLoadBalancerInBound"),
	// 				Name: to.Ptr("AllowAzureLoadBalancerInBound"),
	// 				Properties: &armnetwork.SecurityRulePropertiesFormat{
	// 					Description: to.Ptr("Allow inbound traffic from azure load balancer"),
	// 					Access: to.Ptr(armnetwork.SecurityRuleAccessAllow),
	// 					DestinationAddressPrefix: to.Ptr("*"),
	// 					DestinationPortRange: to.Ptr("*"),
	// 					Direction: to.Ptr(armnetwork.SecurityRuleDirectionInbound),
	// 					Priority: to.Ptr[int32](65001),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					SourceAddressPrefix: to.Ptr("AzureLoadBalancer"),
	// 					SourcePortRange: to.Ptr("*"),
	// 					Protocol: to.Ptr(armnetwork.SecurityRuleProtocolAsterisk),
	// 				},
	// 			},
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/testnsg/defaultSecurityRules/DenyAllInBound"),
	// 				Name: to.Ptr("DenyAllInBound"),
	// 				Properties: &armnetwork.SecurityRulePropertiesFormat{
	// 					Description: to.Ptr("Deny all inbound traffic"),
	// 					Access: to.Ptr(armnetwork.SecurityRuleAccessDeny),
	// 					DestinationAddressPrefix: to.Ptr("*"),
	// 					DestinationPortRange: to.Ptr("*"),
	// 					Direction: to.Ptr(armnetwork.SecurityRuleDirectionInbound),
	// 					Priority: to.Ptr[int32](65500),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					SourceAddressPrefix: to.Ptr("*"),
	// 					SourcePortRange: to.Ptr("*"),
	// 					Protocol: to.Ptr(armnetwork.SecurityRuleProtocolAsterisk),
	// 				},
	// 			},
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/testnsg/defaultSecurityRules/AllowVnetOutBound"),
	// 				Name: to.Ptr("AllowVnetOutBound"),
	// 				Properties: &armnetwork.SecurityRulePropertiesFormat{
	// 					Description: to.Ptr("Allow outbound traffic from all VMs to all VMs in VNET"),
	// 					Access: to.Ptr(armnetwork.SecurityRuleAccessAllow),
	// 					DestinationAddressPrefix: to.Ptr("VirtualNetwork"),
	// 					DestinationPortRange: to.Ptr("*"),
	// 					Direction: to.Ptr(armnetwork.SecurityRuleDirectionOutbound),
	// 					Priority: to.Ptr[int32](65000),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					SourceAddressPrefix: to.Ptr("VirtualNetwork"),
	// 					SourcePortRange: to.Ptr("*"),
	// 					Protocol: to.Ptr(armnetwork.SecurityRuleProtocolAsterisk),
	// 				},
	// 			},
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/testnsg/defaultSecurityRules/AllowInternetOutBound"),
	// 				Name: to.Ptr("AllowInternetOutBound"),
	// 				Properties: &armnetwork.SecurityRulePropertiesFormat{
	// 					Description: to.Ptr("Allow outbound traffic from all VMs to Internet"),
	// 					Access: to.Ptr(armnetwork.SecurityRuleAccessAllow),
	// 					DestinationAddressPrefix: to.Ptr("Internet"),
	// 					DestinationPortRange: to.Ptr("*"),
	// 					Direction: to.Ptr(armnetwork.SecurityRuleDirectionOutbound),
	// 					Priority: to.Ptr[int32](65001),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					SourceAddressPrefix: to.Ptr("*"),
	// 					SourcePortRange: to.Ptr("*"),
	// 					Protocol: to.Ptr(armnetwork.SecurityRuleProtocolAsterisk),
	// 				},
	// 			},
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/testnsg/defaultSecurityRules/DenyAllOutBound"),
	// 				Name: to.Ptr("DenyAllOutBound"),
	// 				Properties: &armnetwork.SecurityRulePropertiesFormat{
	// 					Description: to.Ptr("Deny all outbound traffic"),
	// 					Access: to.Ptr(armnetwork.SecurityRuleAccessDeny),
	// 					DestinationAddressPrefix: to.Ptr("*"),
	// 					DestinationPortRange: to.Ptr("*"),
	// 					Direction: to.Ptr(armnetwork.SecurityRuleDirectionOutbound),
	// 					Priority: to.Ptr[int32](65500),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					SourceAddressPrefix: to.Ptr("*"),
	// 					SourcePortRange: to.Ptr("*"),
	// 					Protocol: to.Ptr(armnetwork.SecurityRuleProtocolAsterisk),
	// 				},
	// 		}},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		SecurityRules: []*armnetwork.SecurityRule{
	// 		},
	// 	},
	// }
}
