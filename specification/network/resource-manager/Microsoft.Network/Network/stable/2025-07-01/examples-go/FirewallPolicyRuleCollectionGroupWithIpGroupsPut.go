package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/FirewallPolicyRuleCollectionGroupWithIpGroupsPut.json
func ExampleFirewallPolicyRuleCollectionGroupsClient_BeginCreateOrUpdate_createFirewallPolicyRuleCollectionGroupWithIPGroups() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFirewallPolicyRuleCollectionGroupsClient().BeginCreateOrUpdate(ctx, "rg1", "firewallPolicy", "ruleCollectionGroup1", armnetwork.FirewallPolicyRuleCollectionGroup{
		Properties: &armnetwork.FirewallPolicyRuleCollectionGroupProperties{
			Priority: to.Ptr[int32](110),
			RuleCollections: []armnetwork.FirewallPolicyRuleCollectionClassification{
				&armnetwork.FirewallPolicyFilterRuleCollection{
					Name: to.Ptr("Example-Filter-Rule-Collection"),
					Action: &armnetwork.FirewallPolicyFilterRuleCollectionAction{
						Type: to.Ptr(armnetwork.FirewallPolicyFilterRuleCollectionActionTypeDeny),
					},
					RuleCollectionType: to.Ptr(armnetwork.FirewallPolicyRuleCollectionTypeFirewallPolicyFilterRuleCollection),
					Rules: []armnetwork.FirewallPolicyRuleClassification{
						&armnetwork.Rule{
							Name: to.Ptr("network-1"),
							DestinationIPGroups: []*string{
								to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/resourceGroups/rg1/ipGroups/ipGroups2"),
							},
							DestinationPorts: []*string{
								to.Ptr("*"),
							},
							IPProtocols: []*armnetwork.FirewallPolicyRuleNetworkProtocol{
								to.Ptr(armnetwork.FirewallPolicyRuleNetworkProtocolTCP),
							},
							RuleType: to.Ptr(armnetwork.FirewallPolicyRuleTypeNetworkRule),
							SourceIPGroups: []*string{
								to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/resourceGroups/rg1/ipGroups/ipGroups1"),
							},
						},
					},
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.FirewallPolicyRuleCollectionGroupsClientCreateOrUpdateResponse{
	// 	FirewallPolicyRuleCollectionGroup: armnetwork.FirewallPolicyRuleCollectionGroup{
	// 		Name: to.Ptr("ruleCollectionGroup1"),
	// 		Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/firewallPolicies/firewallPolicy/ruleCollectionGroups/ruleCollectionGroup1"),
	// 		Properties: &armnetwork.FirewallPolicyRuleCollectionGroupProperties{
	// 			Priority: to.Ptr[int32](110),
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			RuleCollections: []armnetwork.FirewallPolicyRuleCollectionClassification{
	// 				&armnetwork.FirewallPolicyFilterRuleCollection{
	// 					Name: to.Ptr("Example-Filter-Rule-Collection"),
	// 					Action: &armnetwork.FirewallPolicyFilterRuleCollectionAction{
	// 						Type: to.Ptr(armnetwork.FirewallPolicyFilterRuleCollectionActionTypeDeny),
	// 					},
	// 					RuleCollectionType: to.Ptr(armnetwork.FirewallPolicyRuleCollectionTypeFirewallPolicyFilterRuleCollection),
	// 					Rules: []armnetwork.FirewallPolicyRuleClassification{
	// 						&armnetwork.Rule{
	// 							Name: to.Ptr("network-1"),
	// 							DestinationIPGroups: []*string{
	// 								to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/resourceGroups/rg1/ipGroups/ipGroups2"),
	// 							},
	// 							DestinationPorts: []*string{
	// 								to.Ptr("*"),
	// 							},
	// 							IPProtocols: []*armnetwork.FirewallPolicyRuleNetworkProtocol{
	// 								to.Ptr(armnetwork.FirewallPolicyRuleNetworkProtocolTCP),
	// 							},
	// 							RuleType: to.Ptr(armnetwork.FirewallPolicyRuleTypeNetworkRule),
	// 							SourceIPGroups: []*string{
	// 								to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/resourceGroups/rg1/ipGroups/ipGroups1"),
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
