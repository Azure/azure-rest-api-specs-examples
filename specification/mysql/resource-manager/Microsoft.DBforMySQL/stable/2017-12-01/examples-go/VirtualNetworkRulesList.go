package armmysql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/VirtualNetworkRulesList.json
func ExampleVirtualNetworkRulesClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualNetworkRulesClient().NewListByServerPager("TestGroup", "vnet-test-svr", nil)
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
		// page.VirtualNetworkRuleListResult = armmysql.VirtualNetworkRuleListResult{
		// 	Value: []*armmysql.VirtualNetworkRule{
		// 		{
		// 			Name: to.Ptr("vnet-firewall-rule"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/virtualNetworkRules"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforMySQL/servers/vnet-test-svr/virtualNetworkRules/vnet-firewall-rule"),
		// 			Properties: &armmysql.VirtualNetworkRuleProperties{
		// 				IgnoreMissingVnetServiceEndpoint: to.Ptr(false),
		// 				State: to.Ptr(armmysql.VirtualNetworkRuleStateReady),
		// 				VirtualNetworkSubnetID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.Network/virtualNetworks/testvnet/subnets/testsubnet"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("vnet-firewall-rule"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/virtualNetworkRules"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforMySQL/servers/vnet-test-svr/virtualNetworkRules/vnet-firewall-rule"),
		// 			Properties: &armmysql.VirtualNetworkRuleProperties{
		// 				IgnoreMissingVnetServiceEndpoint: to.Ptr(false),
		// 				State: to.Ptr(armmysql.VirtualNetworkRuleStateReady),
		// 				VirtualNetworkSubnetID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.Network/virtualNetworks/testvnet/subnets/testsubnet"),
		// 			},
		// 	}},
		// }
	}
}
