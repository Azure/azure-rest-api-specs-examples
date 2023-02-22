package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetIpFirewallRule.json
func ExampleIPFirewallRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewIPFirewallRulesClient("01234567-89ab-4def-0123-456789abcdef", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "ExampleResourceGroup", "ExampleWorkspace", "ExampleIpFirewallRule", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IPFirewallRuleInfo = armsynapse.IPFirewallRuleInfo{
	// 	Name: to.Ptr("ExampleIpFirewallRule"),
	// 	Type: to.Ptr("Microsoft.Synapse/workspaces/firewallRules"),
	// 	ID: to.Ptr("/subscriptions/01234567-89ab-4def-0123-456789abcdef/resourceGroups/ExampleResourceGroup/providers/Microsoft.Synapse/workspaces/ExampleWorkspace/firewallRules/ExampleIpFirewallRule"),
	// 	Properties: &armsynapse.IPFirewallRuleProperties{
	// 		EndIPAddress: to.Ptr("10.0.0.254"),
	// 		ProvisioningState: to.Ptr(armsynapse.ProvisioningStateSucceeded),
	// 		StartIPAddress: to.Ptr("10.0.0.0"),
	// 	},
	// }
}
