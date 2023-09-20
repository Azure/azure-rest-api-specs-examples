package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/FirewallRuleUpdate.json
func ExampleFirewallRulesClient_CreateOrUpdate_updateAFirewallRuleMaxMin() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFirewallRulesClient().CreateOrUpdate(ctx, "firewallrulecrudtest-12", "firewallrulecrudtest-6285", "firewallrulecrudtest-3927", armsql.FirewallRule{
		Properties: &armsql.ServerFirewallRuleProperties{
			EndIPAddress:   to.Ptr("0.0.0.1"),
			StartIPAddress: to.Ptr("0.0.0.1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FirewallRule = armsql.FirewallRule{
	// 	Name: to.Ptr("firewallrulecrudtest-3927"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/firewallRules"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/firewallrulecrudtest-12/providers/Microsoft.Sql/servers/firewallrulecrudtest-6285/firewallRules/firewallrulecrudtest-3927"),
	// 	Properties: &armsql.ServerFirewallRuleProperties{
	// 		EndIPAddress: to.Ptr("0.0.0.1"),
	// 		StartIPAddress: to.Ptr("0.0.0.1"),
	// 	},
	// }
}
