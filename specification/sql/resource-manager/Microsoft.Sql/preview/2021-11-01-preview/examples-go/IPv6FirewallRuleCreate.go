package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2021-11-01-preview/examples/IPv6FirewallRuleCreate.json
func ExampleIPv6FirewallRulesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsql.NewIPv6FirewallRulesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"firewallrulecrudtest-12",
		"firewallrulecrudtest-6285",
		"firewallrulecrudtest-5370",
		armsql.IPv6FirewallRule{
			Properties: &armsql.IPv6ServerFirewallRuleProperties{
				EndIPv6Address:   to.Ptr("0000:0000:0000:0000:0000:ffff:0000:0003"),
				StartIPv6Address: to.Ptr("0000:0000:0000:0000:0000:ffff:0000:0003"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
