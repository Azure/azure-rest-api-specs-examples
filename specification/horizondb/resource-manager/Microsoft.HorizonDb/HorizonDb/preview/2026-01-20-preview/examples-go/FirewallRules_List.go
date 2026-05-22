package armhorizondb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/horizondb/armhorizondb"
)

// Generated from example definition: 2026-01-20-preview/FirewallRules_List.json
func ExampleFirewallRulesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhorizondb.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFirewallRulesClient().NewListPager("exampleresourcegroup", "examplecluster", "examplepool", nil)
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
		// page = armhorizondb.FirewallRulesClientListResponse{
		// 	FirewallRuleListResult: armhorizondb.FirewallRuleListResult{
		// 		Value: []*armhorizondb.FirewallRule{
		// 			{
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.HorizonDb/clusters/examplecluster/pools/examplepool/firewallRules/allowCorporateNetwork"),
		// 				Name: to.Ptr("allowCorporateNetwork"),
		// 				Type: to.Ptr("Microsoft.HorizonDb/clusters/pools/firewallRules"),
		// 				Properties: &armhorizondb.FirewallRuleProperties{
		// 					StartIPAddress: to.Ptr("10.0.0.1"),
		// 					EndIPAddress: to.Ptr("10.0.0.255"),
		// 					Description: to.Ptr("Allow access from corporate network"),
		// 					ProvisioningState: to.Ptr(armhorizondb.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.HorizonDb/clusters/examplecluster/pools/examplepool/firewallRules/allowDeveloperAccess"),
		// 				Name: to.Ptr("allowDeveloperAccess"),
		// 				Type: to.Ptr("Microsoft.HorizonDb/clusters/pools/firewallRules"),
		// 				Properties: &armhorizondb.FirewallRuleProperties{
		// 					StartIPAddress: to.Ptr("192.168.1.10"),
		// 					EndIPAddress: to.Ptr("192.168.1.10"),
		// 					Description: to.Ptr("Allow developer access"),
		// 					ProvisioningState: to.Ptr(armhorizondb.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
