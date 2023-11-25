package armpolicyinsights_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/05a9cdab363b8ec824094ee73950c04594325172/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_FilterAndAggregateOnly.json
func ExamplePolicyStatesClient_NewListQueryResultsForSubscriptionPager_filterAndAggregateOnly() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPolicyStatesClient().NewListQueryResultsForSubscriptionPager(armpolicyinsights.PolicyStatesResourceLatest, "fffedd8f-ffff-fffd-fffd-fffed2f84852", &armpolicyinsights.QueryOptions{Top: nil,
		Filter:    to.Ptr("PolicyDefinitionAction eq 'deny'"),
		OrderBy:   nil,
		Select:    nil,
		From:      to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-10-05T18:00:00.000Z"); return t }()),
		To:        nil,
		Apply:     to.Ptr("aggregate($count as NumDenyStates)"),
		SkipToken: nil,
		Expand:    nil,
	}, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.PolicyStatesQueryResults = armpolicyinsights.PolicyStatesQueryResults{
		// 	ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/$metadata#latest"),
		// 	ODataCount: to.Ptr[int32](1),
		// 	Value: []*armpolicyinsights.PolicyState{
		// 		{
		// 			AdditionalProperties: map[string]any{
		// 				"NumDenyStates": float64(6),
		// 			},
		// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/$metadata#latest/$entity"),
		// 	}},
		// }
	}
}
