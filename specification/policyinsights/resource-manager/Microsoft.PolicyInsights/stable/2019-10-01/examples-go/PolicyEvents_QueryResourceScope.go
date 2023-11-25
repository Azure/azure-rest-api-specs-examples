package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/05a9cdab363b8ec824094ee73950c04594325172/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyEvents_QueryResourceScope.json
func ExamplePolicyEventsClient_NewListQueryResultsForResourcePager_queryAtResourceScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPolicyEventsClient().NewListQueryResultsForResourcePager(armpolicyinsights.PolicyEventsResourceTypeDefault, "subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourceGroups/myResourceGroup/providers/Microsoft.ClassicCompute/domainNames/myDomainName", &armpolicyinsights.QueryOptions{Top: nil,
		Filter:    nil,
		OrderBy:   nil,
		Select:    nil,
		From:      nil,
		To:        nil,
		Apply:     nil,
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
		// 	ODataContext: to.Ptr("https://management.azure.com/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourcegroups/myResourceGroup/providers/Microsoft.ClassicCompute/domainNames/myDomainName/providers/Microsoft.PolicyInsights/policyEvents/$metadata#default"),
		// 	ODataCount: to.Ptr[int32](2),
		// 	Value: []*armpolicyinsights.PolicyEvent{
		// 		{
		// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourcegroups/myResourceGroup/providers/Microsoft.ClassicCompute/domainNames/myDomainName/providers/Microsoft.PolicyInsights/policyEvents/$metadata#default/$entity"),
		// 			ComplianceState: to.Ptr("NonCompliant"),
		// 			IsCompliant: to.Ptr(false),
		// 			ManagementGroupIDs: to.Ptr("myManagementGroup,fff988bf-fff1-ffff-fffb-fffcd011db47"),
		// 			PolicyAssignmentID: to.Ptr("/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/providers/Microsoft.Authorization/policyAssignments/ec8f9645-8ecb-4abb-9c0b-5292f19d4003"),
		// 			PolicyAssignmentName: to.Ptr("ec8f9645-8ecb-4abb-9c0b-5292f19d4003"),
		// 			PolicyAssignmentOwner: to.Ptr("tbd"),
		// 			PolicyAssignmentParameters: to.Ptr("{}"),
		// 			PolicyAssignmentScope: to.Ptr("/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5"),
		// 			PolicyDefinitionAction: to.Ptr("audit"),
		// 			PolicyDefinitionCategory: to.Ptr("tbd"),
		// 			PolicyDefinitionID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/201ea587-7c90-41c3-910f-c280ae01cfd6"),
		// 			PolicyDefinitionName: to.Ptr("201ea587-7c90-41c3-910f-c280ae01cfd6"),
		// 			PolicySetDefinitionID: to.Ptr("/providers/Microsoft.Authorization/policySetDefinitions/1f3afdf9-d0c9-4c3d-847f-89da613e70a8"),
		// 			PolicySetDefinitionName: to.Ptr("1f3afdf9-d0c9-4c3d-847f-89da613e70a8"),
		// 			PrincipalOid: to.Ptr("fff890fa-fff0-fff3-fff9-fffd7653f078"),
		// 			ResourceGroup: to.Ptr("myResourceGroup"),
		// 			ResourceID: to.Ptr("/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourcegroups/myResourceGroup/providers/Microsoft.ClassicCompute/domainNames/myDomainName"),
		// 			ResourceLocation: to.Ptr("eastus"),
		// 			ResourceTags: to.Ptr("tbd"),
		// 			ResourceType: to.Ptr("/Microsoft.ClassicCompute/domainNames"),
		// 			SubscriptionID: to.Ptr("fff10b27-fff3-fff5-fff8-fffbe01e86a5"),
		// 			TenantID: to.Ptr("fff988bf-fff1-ffff-fffb-fffcd011db47"),
		// 			Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-07T15:14:39.847Z"); return t}()),
		// 		},
		// 		{
		// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourcegroups/myResourceGroup/providers/Microsoft.ClassicCompute/domainNames/myDomainName/providers/Microsoft.PolicyInsights/policyEvents/$metadata#default/$entity"),
		// 			ComplianceState: to.Ptr("NonCompliant"),
		// 			IsCompliant: to.Ptr(false),
		// 			ManagementGroupIDs: to.Ptr("myManagementGroup,fff988bf-fff1-ffff-fffb-fffcd011db47"),
		// 			PolicyAssignmentID: to.Ptr("/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/providers/Microsoft.Authorization/policyAssignments/d15545b8-ff50-409a-a6e3-5bd5cc954003"),
		// 			PolicyAssignmentName: to.Ptr("d15545b8-ff50-409a-a6e3-5bd5cc954003"),
		// 			PolicyAssignmentOwner: to.Ptr("tbd"),
		// 			PolicyAssignmentParameters: to.Ptr("{}"),
		// 			PolicyAssignmentScope: to.Ptr("/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5"),
		// 			PolicyDefinitionAction: to.Ptr("audit"),
		// 			PolicyDefinitionCategory: to.Ptr("tbd"),
		// 			PolicyDefinitionID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/201ea587-7c90-41c3-910f-c280ae01cfd6"),
		// 			PolicyDefinitionName: to.Ptr("201ea587-7c90-41c3-910f-c280ae01cfd6"),
		// 			PolicySetDefinitionID: to.Ptr("/providers/Microsoft.Authorization/policySetDefinitions/1f3afdf9-d0c9-4c3d-847f-89da613e70a8"),
		// 			PolicySetDefinitionName: to.Ptr("1f3afdf9-d0c9-4c3d-847f-89da613e70a8"),
		// 			PrincipalOid: to.Ptr("fff890fa-fff0-fff3-fff9-fffd7653f078"),
		// 			ResourceGroup: to.Ptr("myResourceGroup"),
		// 			ResourceID: to.Ptr("/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourcegroups/myResourceGroup/providers/Microsoft.ClassicCompute/domainNames/myDomainName"),
		// 			ResourceLocation: to.Ptr("eastus"),
		// 			ResourceTags: to.Ptr("tbd"),
		// 			ResourceType: to.Ptr("/Microsoft.ClassicCompute/domainNames"),
		// 			SubscriptionID: to.Ptr("fff10b27-fff3-fff5-fff8-fffbe01e86a5"),
		// 			TenantID: to.Ptr("fff988bf-fff1-ffff-fffb-fffcd011db47"),
		// 			Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-07T15:14:39.584Z"); return t}()),
		// 	}},
		// }
	}
}
