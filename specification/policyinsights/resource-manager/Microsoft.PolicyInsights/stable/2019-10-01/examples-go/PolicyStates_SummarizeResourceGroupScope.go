package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_SummarizeResourceGroupScope.json
func ExamplePolicyStatesClient_SummarizeForResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpolicyinsights.NewPolicyStatesClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.SummarizeForResourceGroup(ctx,
		armpolicyinsights.PolicyStatesSummaryResourceTypeLatest,
		"fffedd8f-ffff-fffd-fffd-fffed2f84852",
		"myResourceGroup",
		&armpolicyinsights.QueryOptions{Top: nil,
			Filter:    nil,
			OrderBy:   nil,
			Select:    nil,
			From:      nil,
			To:        nil,
			Apply:     nil,
			SkipToken: nil,
			Expand:    nil,
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
