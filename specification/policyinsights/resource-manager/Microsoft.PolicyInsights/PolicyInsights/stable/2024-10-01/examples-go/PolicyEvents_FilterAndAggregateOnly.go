package armpolicyinsights_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: 2024-10-01/PolicyEvents_FilterAndAggregateOnly.json
func ExamplePolicyEventsClient_NewListQueryResultsForSubscriptionPager_filterAndAggregateOnly() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPolicyEventsClient().NewListQueryResultsForSubscriptionPager(armpolicyinsights.PolicyEventsResourceTypeDefault, "fffedd8f-ffff-fffd-fffd-fffed2f84852", &armpolicyinsights.PolicyEventsClientListQueryResultsForSubscriptionOptions{
		Apply:  to.Ptr("aggregate($count as NumDenyEvents)"),
		Filter: to.Ptr("PolicyDefinitionAction eq 'deny'"),
		From:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-05T18:00:00Z"); return t }())})
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
		// page = armpolicyinsights.PolicyEventsClientListQueryResultsForSubscriptionResponse{
		// 	PolicyEventsQueryResults: armpolicyinsights.PolicyEventsQueryResults{
		// 		ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyEvents/$metadata#default"),
		// 		ODataCount: to.Ptr[int32](1),
		// 		Value: []*armpolicyinsights.PolicyEvent{
		// 			{
		// 				ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyEvents/$metadata#default/$entity"),
		// 				AdditionalProperties: map[string]any{
		// 				"NumDenyEvents": 40,
		// 			},
		// 			},
		// 		},
		// 	},
		// }
	}
}
