package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/ConfigurationPolicyGroupListByVpnServerConfiguration.json
func ExampleConfigurationPolicyGroupsClient_NewListByVPNServerConfigurationPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConfigurationPolicyGroupsClient().NewListByVPNServerConfigurationPager("rg1", "vpnServerConfiguration1", nil)
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
		// page.ListVPNServerConfigurationPolicyGroupsResult = armnetwork.ListVPNServerConfigurationPolicyGroupsResult{
		// 	Value: []*armnetwork.VPNServerConfigurationPolicyGroup{
		// 		{
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnServerConfigurations/vpnServerConfiguration1/vpnServerConfigurationPolicyGroups/policyGroup1"),
		// 			Name: to.Ptr("policyGroup1"),
		// 			Type: to.Ptr("Microsoft.Network/vpnServerConfigurations/configurationPolicyGroups"),
		// 			Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 			Properties: &armnetwork.VPNServerConfigurationPolicyGroupProperties{
		// 				IsDefault: to.Ptr(true),
		// 				P2SConnectionConfigurations: []*armnetwork.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/p2sVpnGateways/p2sVpnGateway1/p2sConnectionConfigurations/P2SConnectionConfig1"),
		// 				}},
		// 				PolicyMembers: []*armnetwork.VPNServerConfigurationPolicyGroupMember{
		// 					{
		// 						Name: to.Ptr("policy1"),
		// 						AttributeType: to.Ptr(armnetwork.VPNPolicyMemberAttributeTypeRadiusAzureGroupID),
		// 						AttributeValue: to.Ptr("6ad1bd08"),
		// 					},
		// 					{
		// 						Name: to.Ptr("policy2"),
		// 						AttributeType: to.Ptr(armnetwork.VPNPolicyMemberAttributeTypeCertificateGroupID),
		// 						AttributeValue: to.Ptr("red.com"),
		// 				}},
		// 				Priority: to.Ptr[int32](0),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnServerConfigurations/vpnServerConfiguration1/vpnServerConfigurationPolicyGroups/policyGroup2"),
		// 			Name: to.Ptr("policyGroup2"),
		// 			Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 			Properties: &armnetwork.VPNServerConfigurationPolicyGroupProperties{
		// 				IsDefault: to.Ptr(true),
		// 				PolicyMembers: []*armnetwork.VPNServerConfigurationPolicyGroupMember{
		// 					{
		// 						Name: to.Ptr("policy1"),
		// 						AttributeType: to.Ptr(armnetwork.VPNPolicyMemberAttributeTypeRadiusAzureGroupID),
		// 						AttributeValue: to.Ptr("6ad1bd08"),
		// 					},
		// 					{
		// 						Name: to.Ptr("policy2"),
		// 						AttributeType: to.Ptr(armnetwork.VPNPolicyMemberAttributeTypeCertificateGroupID),
		// 						AttributeValue: to.Ptr("red.com"),
		// 				}},
		// 				Priority: to.Ptr[int32](0),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
