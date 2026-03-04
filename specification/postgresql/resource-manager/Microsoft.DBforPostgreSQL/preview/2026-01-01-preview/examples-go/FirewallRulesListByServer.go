package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v6"
)

// Generated from example definition: 2026-01-01-preview/FirewallRulesListByServer.json
func ExampleFirewallRulesClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFirewallRulesClient().NewListByServerPager("exampleresourcegroup", "exampleserver", nil)
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
		// page = armpostgresqlflexibleservers.FirewallRulesClientListByServerResponse{
		// 	FirewallRuleList: armpostgresqlflexibleservers.FirewallRuleList{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff//resourceGroups/exampleresourcegroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/exampleserver/firewallRules?api-version=2025-06-01-preview&$skiptoken=skiptoken"),
		// 		Value: []*armpostgresqlflexibleservers.FirewallRule{
		// 			{
		// 				Name: to.Ptr("examplefirewallrule1"),
		// 				Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/firewallRules"),
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/exampleserver/firewallRules/examplefirewallrule1"),
		// 				Properties: &armpostgresqlflexibleservers.FirewallRuleProperties{
		// 					EndIPAddress: to.Ptr("255.255.255.255"),
		// 					StartIPAddress: to.Ptr("0.0.0.0"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("examplefirewallrule2"),
		// 				Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/firewallRules"),
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/exampleserver/firewallRules/examplefirewallrule2"),
		// 				Properties: &armpostgresqlflexibleservers.FirewallRuleProperties{
		// 					EndIPAddress: to.Ptr("255.0.0.0"),
		// 					StartIPAddress: to.Ptr("1.0.0.0"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
