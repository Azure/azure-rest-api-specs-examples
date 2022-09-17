package armdnsresolver_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dnsresolver/resource-manager/Microsoft.Network/stable/2022-07-01/examples/ForwardingRule_Patch.json
func ExampleForwardingRulesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdnsresolver.NewForwardingRulesClient("abdd4249-9f34-4cc6-8e42-c2e32110603e", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx, "sampleResourceGroup", "sampleDnsForwardingRuleset", "sampleForwardingRule", armdnsresolver.ForwardingRulePatch{
		Properties: &armdnsresolver.ForwardingRulePatchProperties{
			ForwardingRuleState: to.Ptr(armdnsresolver.ForwardingRuleStateDisabled),
			Metadata: map[string]*string{
				"additionalProp2": to.Ptr("value2"),
			},
		},
	}, &armdnsresolver.ForwardingRulesClientUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
