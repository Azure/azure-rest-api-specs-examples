package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/FirewallPolicyRuleCollectionGroupWithWebCategoriesGet.json
func ExampleFirewallPolicyRuleCollectionGroupsClient_Get_getFirewallPolicyRuleCollectionGroupWithWebCategories() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewFirewallPolicyRuleCollectionGroupsClient("e747cc13-97d4-4a79-b463-42d7f4e558f2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "rg1", "firewallPolicy", "ruleCollectionGroup1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
