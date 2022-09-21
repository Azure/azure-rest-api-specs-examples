package armdynatrace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dynatrace/armdynatrace"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2021-09-01/examples/TagRules_Update_MinimumSet_Gen.json
func ExampleTagRulesClient_Update_tagRulesUpdateMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdynatrace.NewTagRulesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx, "myResourceGroup", "myMonitor", "default", armdynatrace.TagRuleUpdate{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
