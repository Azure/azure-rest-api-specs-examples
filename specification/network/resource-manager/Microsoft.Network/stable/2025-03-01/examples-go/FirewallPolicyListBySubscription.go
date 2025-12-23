package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/FirewallPolicyListBySubscription.json
func ExampleFirewallPoliciesClient_NewListAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFirewallPoliciesClient().NewListAllPager(nil)
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
		// page.FirewallPolicyListResult = armnetwork.FirewallPolicyListResult{
		// 	Value: []*armnetwork.FirewallPolicy{
		// 		{
		// 			Name: to.Ptr("firewallPolicy"),
		// 			Type: to.Ptr("Microsoft.Network/firewallPolicies"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/firewallPolicies/firewallPolicy"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 			},
		// 			Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 			Properties: &armnetwork.FirewallPolicyPropertiesFormat{
		// 				DNSSettings: &armnetwork.DNSSettings{
		// 					EnableProxy: to.Ptr(true),
		// 					RequireProxyForNetworkRules: to.Ptr(false),
		// 					Servers: []*string{
		// 						to.Ptr("30.3.4.5")},
		// 					},
		// 					Firewalls: []*armnetwork.SubResource{
		// 					},
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					RuleCollectionGroups: []*armnetwork.SubResource{
		// 						{
		// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/firewallPolicies/firewallPolicy/ruleCollectionGroups/ruleCollectionGroup1"),
		// 					}},
		// 					SKU: &armnetwork.FirewallPolicySKU{
		// 						Tier: to.Ptr(armnetwork.FirewallPolicySKUTierStandard),
		// 					},
		// 					Snat: &armnetwork.FirewallPolicySNAT{
		// 						PrivateRanges: []*string{
		// 							to.Ptr("IANAPrivateRanges")},
		// 						},
		// 						SQL: &armnetwork.FirewallPolicySQL{
		// 							AllowSQLRedirect: to.Ptr(true),
		// 						},
		// 						ThreatIntelMode: to.Ptr(armnetwork.AzureFirewallThreatIntelModeAlert),
		// 					},
		// 			}},
		// 		}
	}
}
