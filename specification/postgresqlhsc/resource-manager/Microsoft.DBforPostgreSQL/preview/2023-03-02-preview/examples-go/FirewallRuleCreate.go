package armcosmosforpostgresql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmosforpostgresql/armcosmosforpostgresql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-03-02-preview/examples/FirewallRuleCreate.json
func ExampleFirewallRulesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmosforpostgresql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFirewallRulesClient().BeginCreateOrUpdate(ctx, "TestGroup", "pgtestsvc4", "rule1", armcosmosforpostgresql.FirewallRule{
		Properties: &armcosmosforpostgresql.FirewallRuleProperties{
			EndIPAddress:   to.Ptr("255.255.255.255"),
			StartIPAddress: to.Ptr("0.0.0.0"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FirewallRule = armcosmosforpostgresql.FirewallRule{
	// 	Name: to.Ptr("rule1"),
	// 	Type: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/firewallRules"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/pgtestsvc4/firewallRules/rule1"),
	// 	SystemData: &armcosmosforpostgresql.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armcosmosforpostgresql.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armcosmosforpostgresql.CreatedByTypeUser),
	// 	},
	// 	Properties: &armcosmosforpostgresql.FirewallRuleProperties{
	// 		EndIPAddress: to.Ptr("255.255.255.255"),
	// 		ProvisioningState: to.Ptr(armcosmosforpostgresql.ProvisioningStateSucceeded),
	// 		StartIPAddress: to.Ptr("0.0.0.0"),
	// 	},
	// }
}
