package armpolicyinsights_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/05a9cdab363b8ec824094ee73950c04594325172/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyEvents_FilterAndMultipleGroups.json
func ExamplePolicyEventsClient_NewListQueryResultsForSubscriptionPager_filterAndMultipleGroups() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPolicyEventsClient().NewListQueryResultsForSubscriptionPager(armpolicyinsights.PolicyEventsResourceTypeDefault, "fffedd8f-ffff-fffd-fffd-fffed2f84852", &armpolicyinsights.QueryOptions{Top: to.Ptr[int32](10),
		Filter:    to.Ptr("PolicyDefinitionAction eq 'deny'"),
		OrderBy:   to.Ptr("NumDeniedResources desc"),
		Select:    nil,
		From:      to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-01T00:00:00.000Z"); return t }()),
		To:        nil,
		Apply:     to.Ptr("groupby((PolicyAssignmentId, PolicyDefinitionId, ResourceId))/groupby((PolicyAssignmentId, PolicyDefinitionId), aggregate($count as NumDeniedResources))"),
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
		// page.PolicyEventsQueryResults = armpolicyinsights.PolicyEventsQueryResults{
		// 	ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyEvents/$metadata#default"),
		// 	ODataCount: to.Ptr[int32](6),
		// 	Value: []*armpolicyinsights.PolicyEvent{
		// 		{
		// 			AdditionalProperties: map[string]any{
		// 				"NumDeniedResources": float64(3),
		// 			},
		// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyEvents/$metadata#default/$entity"),
		// 			PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourcegroups/myResourceGroup1/providers/microsoft.authorization/policyassignments/0591f497c35344fcbaf7a393"),
		// 			PolicyDefinitionID: to.Ptr("/providers/microsoft.authorization/policydefinitions/e56962a6-4747-49cd-b67b-bf8b01975c4c"),
		// 		},
		// 		{
		// 			AdditionalProperties: map[string]any{
		// 				"NumDeniedResources": float64(2),
		// 			},
		// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyEvents/$metadata#default/$entity"),
		// 			PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourcegroups/myResourceGroup1/providers/microsoft.authorization/policyassignments/myassignment1"),
		// 			PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/mydefinition1"),
		// 		},
		// 		{
		// 			AdditionalProperties: map[string]any{
		// 				"NumDeniedResources": float64(2),
		// 			},
		// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyEvents/$metadata#default/$entity"),
		// 			PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourcegroups/myResourceGroup1/providers/microsoft.authorization/policyassignments/myassignment1"),
		// 			PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/mydefinition2"),
		// 		},
		// 		{
		// 			AdditionalProperties: map[string]any{
		// 				"NumDeniedResources": float64(1),
		// 			},
		// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyEvents/$metadata#default/$entity"),
		// 			PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourcegroups/5bc427ca-0089-4d0d-85bd-e98d1e40b3bf/providers/microsoft.authorization/policyassignments/storageaccountsku"),
		// 			PolicyDefinitionID: to.Ptr("/providers/microsoft.authorization/policydefinitions/7433c107-6db4-4ad1-b57a-a76dce0154a1"),
		// 		},
		// 		{
		// 			AdditionalProperties: map[string]any{
		// 				"NumDeniedResources": float64(1),
		// 			},
		// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyEvents/$metadata#default/$entity"),
		// 			PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourcegroups/myResourceGroup2/providers/microsoft.authorization/policyassignments/da43b50031bf4bce84584faa"),
		// 			PolicyDefinitionID: to.Ptr("/providers/microsoft.authorization/policydefinitions/1e30110a-5ceb-460c-a204-c1c3969c6d62"),
		// 		},
		// 		{
		// 			AdditionalProperties: map[string]any{
		// 				"NumDeniedResources": float64(1),
		// 			},
		// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyEvents/$metadata#default/$entity"),
		// 			PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourcegroups/myResourceGroup2/providers/microsoft.authorization/policyassignments/myassignment2"),
		// 			PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/mydefinition3"),
		// 	}},
		// }
	}
}
