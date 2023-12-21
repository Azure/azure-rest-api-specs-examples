package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/FirewallRuleList.json
func ExampleFirewallRulesClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFirewallRulesClient().NewListByServerPager("firewallrulecrudtest-12", "firewallrulecrudtest-6285", nil)
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
		// page.FirewallRuleListResult = armsql.FirewallRuleListResult{
		// 	Value: []*armsql.FirewallRule{
		// 		{
		// 			Name: to.Ptr("firewallrulecrudtest-2304"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/firewallRules"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/firewallrulecrudtest-12/providers/Microsoft.Sql/servers/firewallrulecrudtest-6285/firewallRules/firewallrulecrudtest-2304"),
		// 			Properties: &armsql.ServerFirewallRuleProperties{
		// 				EndIPAddress: to.Ptr("0.0.0.0"),
		// 				StartIPAddress: to.Ptr("0.0.0.0"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("firewallrulecrudtest-3927"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/firewallRules"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/firewallrulecrudtest-12/providers/Microsoft.Sql/servers/firewallrulecrudtest-6285/firewallRules/firewallrulecrudtest-3927"),
		// 			Properties: &armsql.ServerFirewallRuleProperties{
		// 				EndIPAddress: to.Ptr("0.0.0.1"),
		// 				StartIPAddress: to.Ptr("0.0.0.1"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("firewallrulecrudtest-5370"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/firewallRules"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/firewallrulecrudtest-12/providers/Microsoft.Sql/servers/firewallrulecrudtest-6285/firewallRules/firewallrulecrudtest-5370"),
		// 			Properties: &armsql.ServerFirewallRuleProperties{
		// 				EndIPAddress: to.Ptr("0.0.0.3"),
		// 				StartIPAddress: to.Ptr("0.0.0.3"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("firewallrulecrudtest-5767"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/firewallRules"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/firewallrulecrudtest-12/providers/Microsoft.Sql/servers/firewallrulecrudtest-6285/firewallRules/firewallrulecrudtest-5767"),
		// 			Properties: &armsql.ServerFirewallRuleProperties{
		// 				EndIPAddress: to.Ptr("0.0.0.2"),
		// 				StartIPAddress: to.Ptr("0.0.0.2"),
		// 			},
		// 	}},
		// }
	}
}
