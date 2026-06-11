package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/FirewallPolicyRuleCollectionGroupDraftPut.json
func ExampleFirewallPolicyRuleCollectionGroupDraftsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFirewallPolicyRuleCollectionGroupDraftsClient().CreateOrUpdate(ctx, "rg1", "firewallPolicy", "ruleCollectionGroup1", armnetwork.FirewallPolicyRuleCollectionGroupDraft{
		Properties: &armnetwork.FirewallPolicyRuleCollectionGroupDraftProperties{
			Priority: to.Ptr[int32](100),
			RuleCollections: []armnetwork.FirewallPolicyRuleCollectionClassification{
				&armnetwork.FirewallPolicyFilterRuleCollection{
					Name: to.Ptr("Example-Filter-Rule-Collection"),
					Action: &armnetwork.FirewallPolicyFilterRuleCollectionAction{
						Type: to.Ptr(armnetwork.FirewallPolicyFilterRuleCollectionActionTypeDeny),
					},
					Priority:           to.Ptr[int32](100),
					RuleCollectionType: to.Ptr(armnetwork.FirewallPolicyRuleCollectionTypeFirewallPolicyFilterRuleCollection),
					Rules: []armnetwork.FirewallPolicyRuleClassification{
						&armnetwork.Rule{
							Name: to.Ptr("network-rule1"),
							DestinationAddresses: []*string{
								to.Ptr("*"),
							},
							DestinationPorts: []*string{
								to.Ptr("*"),
							},
							IPProtocols: []*armnetwork.FirewallPolicyRuleNetworkProtocol{
								to.Ptr(armnetwork.FirewallPolicyRuleNetworkProtocolTCP),
							},
							RuleType: to.Ptr(armnetwork.FirewallPolicyRuleTypeNetworkRule),
							SourceAddresses: []*string{
								to.Ptr("10.1.25.0/24"),
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
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.FirewallPolicyRuleCollectionGroupDraftsClientCreateOrUpdateResponse{
	// 	FirewallPolicyRuleCollectionGroupDraft: armnetwork.FirewallPolicyRuleCollectionGroupDraft{
	// 		Name: to.Ptr("ruleCollectionGroup1"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/firewallPolicies/firewallPolicy/ruleCollectionGroups/ruleCollectionGroup1"),
	// 		Properties: &armnetwork.FirewallPolicyRuleCollectionGroupDraftProperties{
	// 			Priority: to.Ptr[int32](100),
	// 			RuleCollections: []armnetwork.FirewallPolicyRuleCollectionClassification{
	// 				&armnetwork.FirewallPolicyFilterRuleCollection{
	// 					Name: to.Ptr("Example-Filter-Rule-Collection"),
	// 					Action: &armnetwork.FirewallPolicyFilterRuleCollectionAction{
	// 						Type: to.Ptr(armnetwork.FirewallPolicyFilterRuleCollectionActionTypeDeny),
	// 					},
	// 					Priority: to.Ptr[int32](100),
	// 					RuleCollectionType: to.Ptr(armnetwork.FirewallPolicyRuleCollectionTypeFirewallPolicyFilterRuleCollection),
	// 					Rules: []armnetwork.FirewallPolicyRuleClassification{
	// 						&armnetwork.Rule{
	// 							Name: to.Ptr("network-rule1"),
	// 							DestinationAddresses: []*string{
	// 								to.Ptr("*"),
	// 							},
	// 							DestinationPorts: []*string{
	// 								to.Ptr("*"),
	// 							},
	// 							IPProtocols: []*armnetwork.FirewallPolicyRuleNetworkProtocol{
	// 								to.Ptr(armnetwork.FirewallPolicyRuleNetworkProtocolTCP),
	// 							},
	// 							RuleType: to.Ptr(armnetwork.FirewallPolicyRuleTypeNetworkRule),
	// 							SourceAddresses: []*string{
	// 								to.Ptr("10.1.25.0/24"),
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
