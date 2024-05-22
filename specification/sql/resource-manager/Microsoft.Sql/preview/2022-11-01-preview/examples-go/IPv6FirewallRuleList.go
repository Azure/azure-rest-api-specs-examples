package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/sql/resource-manager/Microsoft.Sql/preview/2022-11-01-preview/examples/IPv6FirewallRuleList.json
func ExampleIPv6FirewallRulesClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIPv6FirewallRulesClient().NewListByServerPager("firewallrulecrudtest-12", "firewallrulecrudtest-6285", nil)
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
		// page.IPv6FirewallRuleListResult = armsql.IPv6FirewallRuleListResult{
		// 	Value: []*armsql.IPv6FirewallRule{
		// 		{
		// 			Name: to.Ptr("firewallrulecrudtest-2304"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/ipv6FirewallRules"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/firewallrulecrudtest-12/providers/Microsoft.Sql/servers/firewallrulecrudtest-6285/ipv6FirewallRules/firewallrulecrudtest-2304"),
		// 			Properties: &armsql.IPv6ServerFirewallRuleProperties{
		// 				EndIPv6Address: to.Ptr("0000:0000:0000:0000:0000:ffff:0000:0000"),
		// 				StartIPv6Address: to.Ptr("0000:0000:0000:0000:0000:ffff:0000:0000"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("firewallrulecrudtest-3927"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/ipv6FirewallRules"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/firewallrulecrudtest-12/providers/Microsoft.Sql/servers/firewallrulecrudtest-6285/ipv6FirewallRules/firewallrulecrudtest-3927"),
		// 			Properties: &armsql.IPv6ServerFirewallRuleProperties{
		// 				EndIPv6Address: to.Ptr("0000:0000:0000:0000:0000:ffff:0000:0001"),
		// 				StartIPv6Address: to.Ptr("0000:0000:0000:0000:0000:ffff:0000:0001"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("firewallrulecrudtest-5370"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/ipv6FirewallRules"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/firewallrulecrudtest-12/providers/Microsoft.Sql/servers/firewallrulecrudtest-6285/ipv6FirewallRules/firewallrulecrudtest-5370"),
		// 			Properties: &armsql.IPv6ServerFirewallRuleProperties{
		// 				EndIPv6Address: to.Ptr("0000:0000:0000:0000:0000:ffff:0000:0003"),
		// 				StartIPv6Address: to.Ptr("0000:0000:0000:0000:0000:ffff:0000:0003"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("firewallrulecrudtest-5767"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/ipv6FirewallRules"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/firewallrulecrudtest-12/providers/Microsoft.Sql/servers/firewallrulecrudtest-6285/ipv6FirewallRules/firewallrulecrudtest-5767"),
		// 			Properties: &armsql.IPv6ServerFirewallRuleProperties{
		// 				EndIPv6Address: to.Ptr("0000:0000:0000:0000:0000:ffff:0000:0002"),
		// 				StartIPv6Address: to.Ptr("0000:0000:0000:0000:0000:ffff:0000:0002"),
		// 			},
		// 	}},
		// }
	}
}
