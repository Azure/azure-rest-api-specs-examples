package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/05a9cdab363b8ec824094ee73950c04594325172/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_FilterAndMultipleGroups.json
func ExamplePolicyStatesClient_NewListQueryResultsForSubscriptionPager_filterAndMultipleGroups() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPolicyStatesClient().NewListQueryResultsForSubscriptionPager(armpolicyinsights.PolicyStatesResourceLatest, "fffedd8f-ffff-fffd-fffd-fffed2f84852", &armpolicyinsights.QueryOptions{Top: to.Ptr[int32](10),
		Filter:    to.Ptr("IsCompliant eq false"),
		OrderBy:   to.Ptr("NumNonCompliantResources desc"),
		Select:    nil,
		From:      nil,
		To:        nil,
		Apply:     to.Ptr("groupby((PolicyAssignmentId, PolicySetDefinitionId, PolicyDefinitionId, PolicyDefinitionReferenceId, ResourceId))/groupby((PolicyAssignmentId, PolicySetDefinitionId, PolicyDefinitionId, PolicyDefinitionReferenceId), aggregate($count as NumNonCompliantResources))"),
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
		// 	ODataCount: to.Ptr[int32](10),
		// 	Value: []*armpolicyinsights.PolicyState{
		// 		{
		// 			AdditionalProperties: map[string]any{
		// 				"NumNonCompliantResources": float64(557),
		// 			},
		// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/$metadata#latest/$entity"),
		// 			PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/f4cc58b7db524a9799381531"),
		// 			PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/24813039-7534-408a-9842-eb99f45721b1"),
		// 			PolicyDefinitionReferenceID: to.Ptr("14799174781370023846"),
		// 			PolicySetDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policysetdefinitions/12b58873-e0f8-4b95-936c-86cbe7c9d697"),
		// 		},
		// 		{
		// 			AdditionalProperties: map[string]any{
		// 				"NumNonCompliantResources": float64(557),
		// 			},
		// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/$metadata#latest/$entity"),
		// 			PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/f4cc58b7db524a9799381531"),
		// 			PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/24813039-7534-408a-9842-eb99f45721b1"),
		// 			PolicyDefinitionReferenceID: to.Ptr("1679708035638239273"),
		// 			PolicySetDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policysetdefinitions/12b58873-e0f8-4b95-936c-86cbe7c9d697"),
		// 		},
		// 		{
		// 			AdditionalProperties: map[string]any{
		// 				"NumNonCompliantResources": float64(557),
		// 			},
		// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/$metadata#latest/$entity"),
		// 			PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/ddd8ef92e3714a5ea3d208c1"),
		// 			PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/24813039-7534-408a-9842-eb99f45721b1"),
		// 			PolicyDefinitionReferenceID: to.Ptr("14799174781370023846"),
		// 			PolicySetDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policysetdefinitions/12b58873-e0f8-4b95-936c-86cbe7c9d697"),
		// 		},
		// 		{
		// 			AdditionalProperties: map[string]any{
		// 				"NumNonCompliantResources": float64(557),
		// 			},
		// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/$metadata#latest/$entity"),
		// 			PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/ddd8ef92e3714a5ea3d208c1"),
		// 			PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/24813039-7534-408a-9842-eb99f45721b1"),
		// 			PolicyDefinitionReferenceID: to.Ptr("1679708035638239273"),
		// 			PolicySetDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policysetdefinitions/12b58873-e0f8-4b95-936c-86cbe7c9d697"),
		// 		},
		// 		{
		// 			AdditionalProperties: map[string]any{
		// 				"NumNonCompliantResources": float64(557),
		// 			},
		// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/$metadata#latest/$entity"),
		// 			PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/d9da7e80af6344ab9d342aa7"),
		// 			PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/24813039-7534-408a-9842-eb99f45721b1"),
		// 			PolicyDefinitionReferenceID: to.Ptr(""),
		// 			PolicySetDefinitionID: to.Ptr(""),
		// 		},
		// 		{
		// 			AdditionalProperties: map[string]any{
		// 				"NumNonCompliantResources": float64(557),
		// 			},
		// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/$metadata#latest/$entity"),
		// 			PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/dcda79d769674aea8bfcaa49"),
		// 			PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/c8b79b49-a579-4045-984e-1b249ab8b474"),
		// 			PolicyDefinitionReferenceID: to.Ptr(""),
		// 			PolicySetDefinitionID: to.Ptr(""),
		// 		},
		// 		{
		// 			AdditionalProperties: map[string]any{
		// 				"NumNonCompliantResources": float64(552),
		// 			},
		// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/$metadata#latest/$entity"),
		// 			PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/3f3c4330183b4e218fe6fd29"),
		// 			PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/24813039-7534-408a-9842-eb99f45721b1"),
		// 			PolicyDefinitionReferenceID: to.Ptr(""),
		// 			PolicySetDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policysetdefinitions/3e3807c1-65c9-49e0-a406-82d8ae3e338c"),
		// 		},
		// 		{
		// 			AdditionalProperties: map[string]any{
		// 				"NumNonCompliantResources": float64(544),
		// 			},
		// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/$metadata#latest/$entity"),
		// 			PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/e46af646ebdb461dba708e01"),
		// 			PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/c8b79b49-a579-4045-984e-1b249ab8b474"),
		// 			PolicyDefinitionReferenceID: to.Ptr("8935913113203900114"),
		// 			PolicySetDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policysetdefinitions/a03db67e-a286-43c3-9098-b2da83d361ad"),
		// 		},
		// 		{
		// 			AdditionalProperties: map[string]any{
		// 				"NumNonCompliantResources": float64(526),
		// 			},
		// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/$metadata#latest/$entity"),
		// 			PolicyAssignmentID: to.Ptr("/providers/microsoft.management/managementgroups/mymg/providers/microsoft.authorization/policyassignments/8174043a1e2849179635b874"),
		// 			PolicyDefinitionID: to.Ptr("/providers/microsoft.management/managementgroups/mymg/providers/microsoft.authorization/policydefinitions/72c0c41a-c752-4bc0-9c61-0d6adc567066"),
		// 			PolicyDefinitionReferenceID: to.Ptr(""),
		// 			PolicySetDefinitionID: to.Ptr(""),
		// 		},
		// 		{
		// 			AdditionalProperties: map[string]any{
		// 				"NumNonCompliantResources": float64(509),
		// 			},
		// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyStates/$metadata#latest/$entity"),
		// 			PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/ddd8ef92e3714a5ea3d208c1"),
		// 			PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/c8b79b49-a579-4045-984e-1b249ab8b474"),
		// 			PolicyDefinitionReferenceID: to.Ptr("2124621540977569058"),
		// 			PolicySetDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policysetdefinitions/12b58873-e0f8-4b95-936c-86cbe7c9d697"),
		// 	}},
		// }
	}
}
