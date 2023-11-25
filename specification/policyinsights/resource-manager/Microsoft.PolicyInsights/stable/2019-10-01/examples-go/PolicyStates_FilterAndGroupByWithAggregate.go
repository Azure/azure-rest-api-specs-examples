package armpolicyinsights_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/05a9cdab363b8ec824094ee73950c04594325172/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_FilterAndGroupByWithAggregate.json
func ExamplePolicyStatesClient_NewListQueryResultsForSubscriptionPager_filterAndGroupWithAggregate() {
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
		Filter:    to.Ptr("IsCompliant eq false and (PolicyDefinitionAction eq 'audit' or PolicyDefinitionAction eq 'deny')"),
		OrderBy:   to.Ptr("NumAuditDenyNonComplianceRecords desc"),
		Select:    nil,
		From:      to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-10-05T18:00:00.000Z"); return t }()),
		To:        nil,
		Apply:     to.Ptr("groupby((PolicyAssignmentId, PolicyDefinitionId, PolicyDefinitionAction, ResourceId), aggregate($count as NumAuditDenyNonComplianceRecords))"),
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
		// 			AdditionalProperties: map[string]any{
		// 				"NumAuditDenyNonComplianceRecords": float64(10),
		// 			},
		// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/$metadata#latest/$entity"),
		// 			PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/e46af646ebdb461dba708e01"),
		// 			PolicyDefinitionAction: to.Ptr("audit"),
		// 			PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/c8b79b49-a579-4045-984e-1b249ab8b474"),
		// 			ResourceID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourcegroups/myrg/providers/microsoft.classiccompute/domainnames/myDomainName"),
		// 		},
		// 		{
		// 			AdditionalProperties: map[string]any{
		// 				"NumAuditDenyNonComplianceRecords": float64(10),
		// 			},
		// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/$metadata#latest/$entity"),
		// 			PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/ddd8ef92e3714a5ea3d208c1"),
		// 			PolicyDefinitionAction: to.Ptr("audit"),
		// 			PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/24813039-7534-408a-9842-eb99f45721b1"),
		// 			ResourceID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourcegroups/myrg/providers/microsoft.classiccompute/domainnames/myDomainName"),
		// 	}},
		// }
	}
}
