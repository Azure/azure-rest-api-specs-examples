package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ReplaceAllIpFirewallRules.json
func ExampleIPFirewallRulesClient_BeginReplaceAll() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewIPFirewallRulesClient("01234567-89ab-4def-0123-456789abcdef", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginReplaceAll(ctx,
		"ExampleResourceGroup",
		"ExampleWorkspace",
		armsynapse.ReplaceAllIPFirewallRulesRequest{
			IPFirewallRules: map[string]*armsynapse.IPFirewallRuleProperties{
				"AnotherExampleFirewallRule": {
					EndIPAddress:   to.Ptr("10.0.1.254"),
					StartIPAddress: to.Ptr("10.0.1.0"),
				},
				"ExampleFirewallRule": {
					EndIPAddress:   to.Ptr("10.0.0.254"),
					StartIPAddress: to.Ptr("10.0.0.0"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
