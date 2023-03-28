package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/sql/resource-manager/Microsoft.Sql/preview/2021-11-01-preview/examples/IPv6FirewallRuleUpdate.json
func ExampleIPv6FirewallRulesClient_CreateOrUpdate_updateAnIPv6FirewallRuleMaxMin() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIPv6FirewallRulesClient().CreateOrUpdate(ctx, "firewallrulecrudtest-12", "firewallrulecrudtest-6285", "firewallrulecrudtest-3927", armsql.IPv6FirewallRule{
		Properties: &armsql.IPv6ServerFirewallRuleProperties{
			EndIPv6Address:   to.Ptr("0000:0000:0000:0000:0000:ffff:0000:0001"),
			StartIPv6Address: to.Ptr("0000:0000:0000:0000:0000:ffff:0000:0001"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IPv6FirewallRule = armsql.IPv6FirewallRule{
	// 	Name: to.Ptr("firewallrulecrudtest-3927"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/ipv6FirewallRules"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/firewallrulecrudtest-12/providers/Microsoft.Sql/servers/firewallrulecrudtest-6285/ipv6FirewallRules/firewallrulecrudtest-3927"),
	// 	Properties: &armsql.IPv6ServerFirewallRuleProperties{
	// 		EndIPv6Address: to.Ptr("0000:0000:0000:0000:0000:ffff:0000:0001"),
	// 		StartIPv6Address: to.Ptr("0000:0000:0000:0000:0000:ffff:0000:0001"),
	// 	},
	// }
}
