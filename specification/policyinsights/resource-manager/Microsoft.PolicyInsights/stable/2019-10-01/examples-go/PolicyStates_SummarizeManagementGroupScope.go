package armpolicyinsights_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_SummarizeManagementGroupScope.json
func ExamplePolicyStatesClient_SummarizeForManagementGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpolicyinsights.NewPolicyStatesClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.SummarizeForManagementGroup(ctx,
		armpolicyinsights.PolicyStatesSummaryResourceTypeLatest,
		"myManagementGroup",
		&armpolicyinsights.QueryOptions{Top: to.Ptr[int32](0),
			Filter:    to.Ptr("PolicyDefinitionAction eq 'deny' or PolicyDefinitionAction eq 'audit'"),
			OrderBy:   nil,
			Select:    nil,
			From:      to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-10-05T18:00:00Z"); return t }()),
			To:        to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-10-06T18:00:00Z"); return t }()),
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
