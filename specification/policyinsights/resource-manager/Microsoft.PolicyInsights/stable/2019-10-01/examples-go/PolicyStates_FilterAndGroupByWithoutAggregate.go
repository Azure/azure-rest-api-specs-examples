package armpolicyinsights_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/05a9cdab363b8ec824094ee73950c04594325172/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_FilterAndGroupByWithoutAggregate.json
func ExamplePolicyStatesClient_NewListQueryResultsForSubscriptionPager_filterAndGroupWithoutAggregate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPolicyStatesClient().NewListQueryResultsForSubscriptionPager(armpolicyinsights.PolicyStatesResourceLatest, "fffedd8f-ffff-fffd-fffd-fffed2f84852", &armpolicyinsights.QueryOptions{Top: to.Ptr[int32](2),
		Filter:    to.Ptr("IsCompliant eq false and (PolicyDefinitionAction ne 'audit' and PolicyDefinitionAction ne 'append')"),
		OrderBy:   nil,
		Select:    nil,
		From:      to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-10-05T18:00:00.000Z"); return t }()),
		To:        nil,
		Apply:     to.Ptr("groupby((PolicyAssignmentId, PolicyDefinitionId, PolicyDefinitionAction, ResourceId))"),
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
		// 	ODataCount: to.Ptr[int32](2),
		// 	Value: []*armpolicyinsights.PolicyState{
		// 		{
		// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/$metadata#latest/$entity"),
		// 			PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/enable monitoring in azure security center"),
		// 			PolicyDefinitionAction: to.Ptr("auditifnotexists"),
		// 			PolicyDefinitionID: to.Ptr("/providers/microsoft.authorization/policydefinitions/44452482-524f-4bf4-b852-0bff7cc4a3ed"),
		// 			ResourceID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourcegroups/myrg/providers/microsoft.network/virtualnetworks/vnet"),
		// 		},
		// 		{
		// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/$metadata#latest/$entity"),
		// 			PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/89b27f38-e9e4-4468-ab81-801c84b8c017"),
		// 			PolicyDefinitionAction: to.Ptr("auditifnotexists"),
		// 			PolicyDefinitionID: to.Ptr("/providers/microsoft.authorization/policydefinitions/44452482-524f-4bf4-b852-0bff7cc4a3ed"),
		// 			ResourceID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourcegroups/myrg/providers/microsoft.network/virtualnetworks/vnet"),
		// 	}},
		// }
	}
}
