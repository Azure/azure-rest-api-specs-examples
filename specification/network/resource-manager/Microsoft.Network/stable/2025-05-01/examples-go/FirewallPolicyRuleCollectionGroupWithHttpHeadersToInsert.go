package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/FirewallPolicyRuleCollectionGroupWithHttpHeadersToInsert.json
func ExampleFirewallPolicyRuleCollectionGroupsClient_BeginCreateOrUpdate_createFirewallPolicyRuleCollectionGroupWithHttpHeaderToInsert() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFirewallPolicyRuleCollectionGroupsClient().BeginCreateOrUpdate(ctx, "rg1", "firewallPolicy", "ruleCollectionGroup1", armnetwork.FirewallPolicyRuleCollectionGroup{
		Properties: &armnetwork.FirewallPolicyRuleCollectionGroupProperties{
			Priority: to.Ptr[int32](110),
			RuleCollections: []armnetwork.FirewallPolicyRuleCollectionClassification{
				&armnetwork.FirewallPolicyFilterRuleCollection{
					Name:               to.Ptr("Example-Filter-Rule-Collection"),
					RuleCollectionType: to.Ptr(armnetwork.FirewallPolicyRuleCollectionTypeFirewallPolicyFilterRuleCollection),
					Action: &armnetwork.FirewallPolicyFilterRuleCollectionAction{
						Type: to.Ptr(armnetwork.FirewallPolicyFilterRuleCollectionActionTypeAllow),
					},
					Rules: []armnetwork.FirewallPolicyRuleClassification{
						&armnetwork.ApplicationRule{
							Name:        to.Ptr("rule1"),
							Description: to.Ptr("Insert trusted tenants header"),
							RuleType:    to.Ptr(armnetwork.FirewallPolicyRuleTypeApplicationRule),
							FqdnTags: []*string{
								to.Ptr("WindowsVirtualDesktop")},
							HTTPHeadersToInsert: []*armnetwork.FirewallPolicyHTTPHeaderToInsert{
								{
									HeaderName:  to.Ptr("Restrict-Access-To-Tenants"),
									HeaderValue: to.Ptr("contoso.com,fabrikam.onmicrosoft.com"),
								}},
							Protocols: []*armnetwork.FirewallPolicyRuleApplicationProtocol{
								{
									Port:         to.Ptr[int32](80),
									ProtocolType: to.Ptr(armnetwork.FirewallPolicyRuleApplicationProtocolTypeHTTP),
								}},
							SourceAddresses: []*string{
								to.Ptr("216.58.216.164"),
								to.Ptr("10.0.0.0/24")},
						}},
				}},
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
	// res.FirewallPolicyRuleCollectionGroup = armnetwork.FirewallPolicyRuleCollectionGroup{
	// 	ID: to.Ptr("/subscriptions/e747cc13-97d4-4a79-b463-42d7f4e558f2/resourceGroups/rg1/providers/Microsoft.Network/firewallPolicies/firewallPolicy/ruleCollectionGroups/ruleCollectionGroup1"),
	// 	Name: to.Ptr("ruleCollectionGroup1"),
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.FirewallPolicyRuleCollectionGroupProperties{
	// 		Priority: to.Ptr[int32](110),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		RuleCollections: []armnetwork.FirewallPolicyRuleCollectionClassification{
	// 			&armnetwork.FirewallPolicyFilterRuleCollection{
	// 				Name: to.Ptr("Example-Filter-Rule-Collection"),
	// 				RuleCollectionType: to.Ptr(armnetwork.FirewallPolicyRuleCollectionTypeFirewallPolicyFilterRuleCollection),
	// 				Action: &armnetwork.FirewallPolicyFilterRuleCollectionAction{
	// 					Type: to.Ptr(armnetwork.FirewallPolicyFilterRuleCollectionActionTypeAllow),
	// 				},
	// 				Rules: []armnetwork.FirewallPolicyRuleClassification{
	// 					&armnetwork.ApplicationRule{
	// 						Name: to.Ptr("rule1"),
	// 						Description: to.Ptr("Insert trusted tenants header"),
	// 						RuleType: to.Ptr(armnetwork.FirewallPolicyRuleTypeApplicationRule),
	// 						FqdnTags: []*string{
	// 							to.Ptr("WindowsVirtualDesktop")},
	// 							HTTPHeadersToInsert: []*armnetwork.FirewallPolicyHTTPHeaderToInsert{
	// 								{
	// 									HeaderName: to.Ptr("Restrict-Access-To-Tenants"),
	// 									HeaderValue: to.Ptr("contoso.com,fabrikam.onmicrosoft.com"),
	// 							}},
	// 							Protocols: []*armnetwork.FirewallPolicyRuleApplicationProtocol{
	// 								{
	// 									Port: to.Ptr[int32](80),
	// 									ProtocolType: to.Ptr(armnetwork.FirewallPolicyRuleApplicationProtocolTypeHTTP),
	// 							}},
	// 							SourceAddresses: []*string{
	// 								to.Ptr("216.58.216.164"),
	// 								to.Ptr("10.0.0.0/24")},
	// 						}},
	// 				}},
	// 			},
	// 		}
}
