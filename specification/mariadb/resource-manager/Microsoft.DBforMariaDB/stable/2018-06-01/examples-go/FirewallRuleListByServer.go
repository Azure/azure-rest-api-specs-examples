package armmariadb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/FirewallRuleListByServer.json
func ExampleFirewallRulesClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmariadb.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFirewallRulesClient().NewListByServerPager("TestGroup", "testserver", nil)
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
		// page.FirewallRuleListResult = armmariadb.FirewallRuleListResult{
		// 	Value: []*armmariadb.FirewallRule{
		// 		{
		// 			Name: to.Ptr("rule1"),
		// 			Type: to.Ptr("Microsoft.DBforMariaDB/servers/firewallRules"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforMariaDB/servers/testserver/firewallRules/rule1"),
		// 			Properties: &armmariadb.FirewallRuleProperties{
		// 				EndIPAddress: to.Ptr("255.255.255.255"),
		// 				StartIPAddress: to.Ptr("0.0.0.0"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("rule2"),
		// 			Type: to.Ptr("Microsoft.DBforMariaDB/servers/firewallRules"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforMariaDB/servers/testserver/firewallRules/rule2"),
		// 			Properties: &armmariadb.FirewallRuleProperties{
		// 				EndIPAddress: to.Ptr("255.0.0.0"),
		// 				StartIPAddress: to.Ptr("1.0.0.0"),
		// 			},
		// 	}},
		// }
	}
}
