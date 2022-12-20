package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/ConfigurationPolicyGroupPut.json
func ExampleConfigurationPolicyGroupsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewConfigurationPolicyGroupsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "vpnServerConfiguration1", "policyGroup1", armnetwork.VPNServerConfigurationPolicyGroup{
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
	// TODO: use response item
	_ = res
}
