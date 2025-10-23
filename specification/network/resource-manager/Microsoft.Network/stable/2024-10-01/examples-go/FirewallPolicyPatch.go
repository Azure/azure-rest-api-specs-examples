package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a9dbb28e788355a47dc5bad3ea5f8da212b4bf6/specification/network/resource-manager/Microsoft.Network/stable/2024-10-01/examples/FirewallPolicyPatch.json
func ExampleFirewallPoliciesClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFirewallPoliciesClient().UpdateTags(ctx, "myResourceGroup", "firewallPolicy", armnetwork.TagsObject{
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
			"key2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FirewallPolicy = armnetwork.FirewallPolicy{
	// 	Name: to.Ptr("firewallPolicy"),
	// 	Type: to.Ptr("Microsoft.Network/firewallPolicies"),
	// 	ID: to.Ptr("/subscriptions/subId/resourceGroups/myResourceGroup/providers/Microsoft.Network/firewallPolicies/firewallPolicy"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 		"key2": to.Ptr("value2"),
	// 	},
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.FirewallPolicyPropertiesFormat{
	// 		DNSSettings: &armnetwork.DNSSettings{
	// 			EnableProxy: to.Ptr(true),
	// 			RequireProxyForNetworkRules: to.Ptr(false),
	// 			Servers: []*string{
	// 				to.Ptr("30.3.4.5")},
	// 			},
	// 			Firewalls: []*armnetwork.SubResource{
	// 			},
	// 			Insights: &armnetwork.FirewallPolicyInsights{
	// 				IsEnabled: to.Ptr(true),
	// 				LogAnalyticsResources: &armnetwork.FirewallPolicyLogAnalyticsResources{
	// 					DefaultWorkspaceID: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourcegroups/rg1/providers/microsoft.operationalinsights/workspaces/defaultWorkspace"),
	// 					},
	// 					Workspaces: []*armnetwork.FirewallPolicyLogAnalyticsWorkspace{
	// 						{
	// 							Region: to.Ptr("westus"),
	// 							WorkspaceID: &armnetwork.SubResource{
	// 								ID: to.Ptr("/subscriptions/subid/resourcegroups/rg1/providers/microsoft.operationalinsights/workspaces/workspace1"),
	// 							},
	// 						},
	// 						{
	// 							Region: to.Ptr("eastus"),
	// 							WorkspaceID: &armnetwork.SubResource{
	// 								ID: to.Ptr("/subscriptions/subid/resourcegroups/rg1/providers/microsoft.operationalinsights/workspaces/workspace2"),
	// 							},
	// 					}},
	// 				},
	// 				RetentionDays: to.Ptr[int32](100),
	// 			},
	// 			IntrusionDetection: &armnetwork.FirewallPolicyIntrusionDetection{
	// 				Configuration: &armnetwork.FirewallPolicyIntrusionDetectionConfiguration{
	// 					BypassTrafficSettings: []*armnetwork.FirewallPolicyIntrusionDetectionBypassTrafficSpecifications{
	// 						{
	// 							Name: to.Ptr("bypassRule1"),
	// 							Description: to.Ptr("Rule 1"),
	// 							DestinationAddresses: []*string{
	// 								to.Ptr("5.6.7.8")},
	// 								DestinationPorts: []*string{
	// 									to.Ptr("*")},
	// 									SourceAddresses: []*string{
	// 										to.Ptr("1.2.3.4")},
	// 										Protocol: to.Ptr(armnetwork.FirewallPolicyIntrusionDetectionProtocolTCP),
	// 								}},
	// 								SignatureOverrides: []*armnetwork.FirewallPolicyIntrusionDetectionSignatureSpecification{
	// 									{
	// 										ID: to.Ptr("2525004"),
	// 										Mode: to.Ptr(armnetwork.FirewallPolicyIntrusionDetectionStateTypeDeny),
	// 								}},
	// 							},
	// 							Mode: to.Ptr(armnetwork.FirewallPolicyIntrusionDetectionStateTypeAlert),
	// 						},
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 						RuleCollectionGroups: []*armnetwork.SubResource{
	// 							{
	// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/firewallPolicies/firewallPolicy/ruleCollectionGroups/ruleCollectionGroup1"),
	// 						}},
	// 						SKU: &armnetwork.FirewallPolicySKU{
	// 							Tier: to.Ptr(armnetwork.FirewallPolicySKUTierPremium),
	// 						},
	// 						Snat: &armnetwork.FirewallPolicySNAT{
	// 							PrivateRanges: []*string{
	// 								to.Ptr("IANAPrivateRanges")},
	// 							},
	// 							SQL: &armnetwork.FirewallPolicySQL{
	// 								AllowSQLRedirect: to.Ptr(true),
	// 							},
	// 							ThreatIntelMode: to.Ptr(armnetwork.AzureFirewallThreatIntelModeAlert),
	// 							ThreatIntelWhitelist: &armnetwork.FirewallPolicyThreatIntelWhitelist{
	// 								Fqdns: []*string{
	// 									to.Ptr("*.microsoft.com")},
	// 									IPAddresses: []*string{
	// 										to.Ptr("20.3.4.5")},
	// 									},
	// 									TransportSecurity: &armnetwork.FirewallPolicyTransportSecurity{
	// 										CertificateAuthority: &armnetwork.FirewallPolicyCertificateAuthority{
	// 											Name: to.Ptr("clientcert"),
	// 											KeyVaultSecretID: to.Ptr("https://kv/secret"),
	// 										},
	// 									},
	// 								},
	// 							}
}
