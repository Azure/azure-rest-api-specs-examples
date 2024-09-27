package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a3913f4b26467aed413cdc907116e99894f08994/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/VirtualNetworkRulesGet.json
func ExampleVirtualNetworkRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualNetworkRulesClient().Get(ctx, "Default", "vnet-test-svr", "vnet-firewall-rule", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualNetworkRule = armsql.VirtualNetworkRule{
	// 	Name: to.Ptr("vnet-firewall-rule"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/virtualNetworkRules"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/vnet-test-svr/virtualNetworkRules/vnet-firewall-rule"),
	// 	Properties: &armsql.VirtualNetworkRuleProperties{
	// 		IgnoreMissingVnetServiceEndpoint: to.Ptr(false),
	// 		State: to.Ptr(armsql.VirtualNetworkRuleStateReady),
	// 		VirtualNetworkSubnetID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Network/virtualNetworks/testvnet/subnets/testsubnet"),
	// 	},
	// }
}
