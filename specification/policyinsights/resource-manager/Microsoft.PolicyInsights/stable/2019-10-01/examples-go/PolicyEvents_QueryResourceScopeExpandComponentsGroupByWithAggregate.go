package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/05a9cdab363b8ec824094ee73950c04594325172/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyEvents_QueryResourceScopeExpandComponentsGroupByWithAggregate.json
func ExamplePolicyEventsClient_NewListQueryResultsForResourcePager_queryComponentsPolicyEventsCountGroupedByUserAndActionTypeForResourceScopeFilteredByGivenAssignment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPolicyEventsClient().NewListQueryResultsForResourcePager(armpolicyinsights.PolicyEventsResourceTypeDefault, "subscriptions/e78961ba-36fe-4739-9212-e3031b4c8db7/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/Vaults/myKVName", &armpolicyinsights.QueryOptions{Top: nil,
		Filter:    to.Ptr("policyAssignmentId eq '/subscriptions/e78961ba-36fe-4739-9212-e3031b4c8db7/providers/microsoft.authorization/policyassignments/560050f83dbb4a24974323f8'"),
		OrderBy:   nil,
		Select:    nil,
		From:      nil,
		To:        nil,
		Apply:     nil,
		SkipToken: nil,
		Expand:    to.Ptr("components($apply=groupby((tenantId, principalOid, policyDefinitionAction), aggregate($count as totalActions)))"),
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
		// 	ODataContext: to.Ptr("https://management.azure.com/subscriptions/e78961ba-36fe-4739-9212-e3031b4c8db7/resourcegroups/myResourceGroup/providers/microsoft.keyvault/vaults/myKVName/providers/Microsoft.PolicyInsights/policyEvents/$metadata#default"),
		// 	ODataCount: to.Ptr[int32](1),
		// 	Value: []*armpolicyinsights.PolicyEvent{
		// 		{
		// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/e78961ba-36fe-4739-9212-e3031b4c8db7/resourcegroups/myResourceGroup/providers/microsoft.keyvault/vaults/myKVName/providers/Microsoft.PolicyInsights/policyEvents/$metadata#default/$entity"),
		// 			ComplianceState: to.Ptr("NonCompliant"),
		// 			Components: []*armpolicyinsights.ComponentEventDetails{
		// 				{
		// 					AdditionalProperties: map[string]any{
		// 						"totalActions": float64(6),
		// 					},
		// 					PolicyDefinitionAction: to.Ptr("audit"),
		// 					PrincipalOid: to.Ptr("0d81b461-6bb0-4909-a102-d51803a7d275"),
		// 					TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
		// 			}},
		// 			EffectiveParameters: to.Ptr(""),
		// 			IsCompliant: to.Ptr(false),
		// 			ManagementGroupIDs: to.Ptr("PolicyUIMG,AzGovTest5,72f988bf-86f1-41af-91ab-2d7cd011db47"),
		// 			PolicyAssignmentID: to.Ptr("/subscriptions/e78961ba-36fe-4739-9212-e3031b4c8db7/providers/microsoft.authorization/policyassignments/560050f83dbb4a24974323f8"),
		// 			PolicyAssignmentName: to.Ptr("560050f83dbb4a24974323f8"),
		// 			PolicyAssignmentOwner: to.Ptr("tbd"),
		// 			PolicyAssignmentParameters: to.Ptr(""),
		// 			PolicyAssignmentScope: to.Ptr("/subscriptions/e78961ba-36fe-4739-9212-e3031b4c8db7"),
		// 			PolicyDefinitionAction: to.Ptr("audit"),
		// 			PolicyDefinitionCategory: to.Ptr("tbd"),
		// 			PolicyDefinitionID: to.Ptr("/subscriptions/e78961ba-36fe-4739-9212-e3031b4c8db7/providers/microsoft.authorization/policydefinitions/ab108bc4-32df-4677-8b38-fa8b2905df59"),
		// 			PolicyDefinitionName: to.Ptr("ab108bc4-32df-4677-8b38-fa8b2905df59"),
		// 			PolicyDefinitionReferenceID: to.Ptr(""),
		// 			PolicySetDefinitionCategory: to.Ptr(""),
		// 			PolicySetDefinitionID: to.Ptr(""),
		// 			PolicySetDefinitionName: to.Ptr(""),
		// 			PolicySetDefinitionOwner: to.Ptr(""),
		// 			PolicySetDefinitionParameters: to.Ptr(""),
		// 			PrincipalOid: to.Ptr(""),
		// 			ResourceGroup: to.Ptr("myResourceGroup"),
		// 			ResourceID: to.Ptr("/subscriptions/e78961ba-36fe-4739-9212-e3031b4c8db7/resourcegroups/myResourceGroup/providers/microsoft.keyvault/vaults/myKVName"),
		// 			ResourceLocation: to.Ptr("westcentralus"),
		// 			ResourceTags: to.Ptr("tbd"),
		// 			ResourceType: to.Ptr("Microsoft.KeyVault/vaults"),
		// 			SubscriptionID: to.Ptr("e78961ba-36fe-4739-9212-e3031b4c8db7"),
		// 			TenantID: to.Ptr(""),
		// 			Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-21T19:42:08.325Z"); return t}()),
		// 	}},
		// }
	}
}
