package armredis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/45374f48f560b3337ed55735038f1e9bf8cbea65/specification/redis/resource-manager/Microsoft.Cache/stable/2024-11-01/examples/RedisCacheFirewallRuleCreate.json
func ExampleFirewallRulesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFirewallRulesClient().CreateOrUpdate(ctx, "rg1", "cache1", "rule1", armredis.FirewallRule{
		Properties: &armredis.FirewallRuleProperties{
			EndIP:   to.Ptr("192.168.1.4"),
			StartIP: to.Ptr("192.168.1.1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FirewallRule = armredis.FirewallRule{
	// 	Name: to.Ptr("cache1/rule1"),
	// 	Type: to.Ptr("Microsoft.Cache/Redis/firewallRules"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cache/Redis/cache1/firewallRules/rule1"),
	// 	Properties: &armredis.FirewallRuleProperties{
	// 		EndIP: to.Ptr("192.168.1.4"),
	// 		StartIP: to.Ptr("192.168.1.1"),
	// 	},
	// }
}
