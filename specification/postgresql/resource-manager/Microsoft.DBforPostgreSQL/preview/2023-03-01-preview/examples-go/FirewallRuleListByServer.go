package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/66a59c94238bf973680355fb179fade4c9405710/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-03-01-preview/examples/FirewallRuleListByServer.json
func ExampleFirewallRulesClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFirewallRulesClient().NewListByServerPager("testrg", "testserver", nil)
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
		// page.FirewallRuleListResult = armpostgresqlflexibleservers.FirewallRuleListResult{
		// 	Value: []*armpostgresqlflexibleservers.FirewallRule{
		// 		{
		// 			Name: to.Ptr("rule1"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/firewallRules"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/testserver/firewallRules/rule1"),
		// 			Properties: &armpostgresqlflexibleservers.FirewallRuleProperties{
		// 				EndIPAddress: to.Ptr("255.255.255.255"),
		// 				StartIPAddress: to.Ptr("0.0.0.0"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("rule2"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/firewallRules"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/testserver/firewallRules/rule2"),
		// 			Properties: &armpostgresqlflexibleservers.FirewallRuleProperties{
		// 				EndIPAddress: to.Ptr("255.0.0.0"),
		// 				StartIPAddress: to.Ptr("1.0.0.0"),
		// 			},
		// 	}},
		// }
	}
}
