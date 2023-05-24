package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/ConfigurationPolicyGroupPut.json
func ExampleConfigurationPolicyGroupsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewConfigurationPolicyGroupsClient().BeginCreateOrUpdate(ctx, "rg1", "vpnServerConfiguration1", "policyGroup1", armnetwork.VPNServerConfigurationPolicyGroup{
		Properties: &armnetwork.VPNServerConfigurationPolicyGroupProperties{
			IsDefault: to.Ptr(true),
			PolicyMembers: []*armnetwork.VPNServerConfigurationPolicyGroupMember{
				{
					Name:           to.Ptr("policy1"),
					AttributeType:  to.Ptr(armnetwork.VPNPolicyMemberAttributeTypeRadiusAzureGroupID),
					AttributeValue: to.Ptr("6ad1bd08"),
				},
				{
					Name:           to.Ptr("policy2"),
					AttributeType:  to.Ptr(armnetwork.VPNPolicyMemberAttributeTypeCertificateGroupID),
					AttributeValue: to.Ptr("red.com"),
				}},
			Priority: to.Ptr[int32](0),
		},
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
	// res.VPNServerConfigurationPolicyGroup = armnetwork.VPNServerConfigurationPolicyGroup{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnServerConfigurations/vpnServerConfiguration1/vpnServerConfigurationPolicyGroups/policyGroup1"),
	// 	Name: to.Ptr("policyGroup1"),
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.VPNServerConfigurationPolicyGroupProperties{
	// 		IsDefault: to.Ptr(true),
	// 		P2SConnectionConfigurations: []*armnetwork.SubResource{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/p2sVpnGateways/p2sVpnGateway1/p2sConnectionConfigurations/P2SConnectionConfig1"),
	// 		}},
	// 		PolicyMembers: []*armnetwork.VPNServerConfigurationPolicyGroupMember{
	// 			{
	// 				Name: to.Ptr("policy1"),
	// 				AttributeType: to.Ptr(armnetwork.VPNPolicyMemberAttributeTypeRadiusAzureGroupID),
	// 				AttributeValue: to.Ptr("6ad1bd08"),
	// 			},
	// 			{
	// 				Name: to.Ptr("policy2"),
	// 				AttributeType: to.Ptr(armnetwork.VPNPolicyMemberAttributeTypeCertificateGroupID),
	// 				AttributeValue: to.Ptr("red.com"),
	// 		}},
	// 		Priority: to.Ptr[int32](0),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 	},
	// }
}
