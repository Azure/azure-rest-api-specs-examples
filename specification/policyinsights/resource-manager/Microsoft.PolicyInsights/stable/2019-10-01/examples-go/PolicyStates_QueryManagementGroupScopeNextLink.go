package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/05a9cdab363b8ec824094ee73950c04594325172/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_QueryManagementGroupScopeNextLink.json
func ExamplePolicyStatesClient_NewListQueryResultsForManagementGroupPager_queryLatestAtManagementGroupScopeWithNextLink() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPolicyStatesClient().NewListQueryResultsForManagementGroupPager(armpolicyinsights.PolicyStatesResourceLatest, "myManagementGroup", &armpolicyinsights.QueryOptions{Top: nil,
		Filter:    nil,
		OrderBy:   nil,
		Select:    nil,
		From:      nil,
		To:        nil,
		Apply:     nil,
		SkipToken: to.Ptr("WpmWfBSvPhkAK6QD"),
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
		// 	ODataContext: to.Ptr("https://management.azure.com/providers/Microsoft.Management/managementGroups/myManagementGroup/providers/Microsoft.PolicyInsights/policyStates/$metadata#latest"),
		// 	ODataCount: to.Ptr[int32](2),
		// 	Value: []*armpolicyinsights.PolicyState{
		// 		{
		// 			ODataContext: to.Ptr("https://management.azure.com/providers/Microsoft.Management/managementGroups/myManagementGroup/providers/Microsoft.PolicyInsights/policyStates/$metadata#latest/$entity"),
		// 			ComplianceState: to.Ptr("NonCompliant"),
		// 			IsCompliant: to.Ptr(false),
		// 			ManagementGroupIDs: to.Ptr("myManagementGroup,fff988bf-fff1-ffff-fffb-fffcd011db47"),
		// 			PolicyAssignmentID: to.Ptr("/providers/Microsoft.Management/managementGroups/myManagementGroup/providers/Microsoft.Authorization/policyAssignments/1ef5d536aec743a0aa801c1a"),
		// 			PolicyAssignmentName: to.Ptr("1ef5d536aec743a0aa801c1a"),
		// 			PolicyAssignmentOwner: to.Ptr("tbd"),
		// 			PolicyAssignmentParameters: to.Ptr("{}"),
		// 			PolicyAssignmentScope: to.Ptr("/providers/Microsoft.Management/managementGroups/myManagementGroup"),
		// 			PolicyDefinitionAction: to.Ptr("audit"),
		// 			PolicyDefinitionCategory: to.Ptr("tbd"),
		// 			PolicyDefinitionGroupNames: []*string{
		// 				to.Ptr("myGroup")},
		// 				PolicyDefinitionID: to.Ptr("/providers/Microsoft.Management/managementGroups/myManagementGroup/providers/Microsoft.Authorization/policyDefinitions/022d9357-5a90-46f7-9554-21d30ce4c32d"),
		// 				PolicyDefinitionName: to.Ptr("022d9357-5a90-46f7-9554-21d30ce4c32d"),
		// 				PolicyDefinitionReferenceID: to.Ptr("15521232277412542086"),
		// 				PolicySetDefinitionID: to.Ptr("/providers/Microsoft.Management/managementGroups/myManagementGroup/providers/Microsoft.Authorization/policySetDefinitions/335cefd2-ab16-430f-b364-974a170eb1d5"),
		// 				PolicySetDefinitionName: to.Ptr("335cefd2-ab16-430f-b364-974a170eb1d5"),
		// 				ResourceGroup: to.Ptr("myrg1"),
		// 				ResourceID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourceGroups/myrg1/providers/microsoft.insights/autoscalesettings/mytest1"),
		// 				ResourceLocation: to.Ptr("westus"),
		// 				ResourceTags: to.Ptr("tbd"),
		// 				ResourceType: to.Ptr("/microsoft.insights/autoscalesettings"),
		// 				SubscriptionID: to.Ptr("fffedd8f-ffff-fffd-fffd-fffed2f84852"),
		// 				Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-10-09T17:41:47.000Z"); return t}()),
		// 			},
		// 			{
		// 				ODataContext: to.Ptr("https://management.azure.com/providers/Microsoft.Management/managementGroups/myManagementGroup/providers/Microsoft.PolicyInsights/policyStates/$metadata#latest/$entity"),
		// 				ComplianceState: to.Ptr("Compliant"),
		// 				IsCompliant: to.Ptr(true),
		// 				ManagementGroupIDs: to.Ptr("myManagementGroup,fff988bf-fff1-ffff-fffb-fffcd011db47"),
		// 				PolicyAssignmentID: to.Ptr("/providers/Microsoft.Management/managementGroups/myManagementGroup/providers/Microsoft.Authorization/policyAssignments/186044306c044a1d8c0ff76c"),
		// 				PolicyAssignmentName: to.Ptr("186044306c044a1d8c0ff76c"),
		// 				PolicyAssignmentOwner: to.Ptr("tbd"),
		// 				PolicyAssignmentParameters: to.Ptr("{\"allowedLocations\":{\"value\":[\"centralus\"]}}"),
		// 				PolicyAssignmentScope: to.Ptr("/providers/Microsoft.Management/managementGroups/myManagementGroup"),
		// 				PolicyDefinitionAction: to.Ptr("audit"),
		// 				PolicyDefinitionCategory: to.Ptr("tbd"),
		// 				PolicyDefinitionGroupNames: []*string{
		// 					to.Ptr("myGroup")},
		// 					PolicyDefinitionID: to.Ptr("/providers/Microsoft.Management/managementGroups/myManagementGroup/providers/Microsoft.Authorization/policyDefinitions/022d9357-5a90-46f7-9554-21d30ce4c32d"),
		// 					PolicyDefinitionName: to.Ptr("022d9357-5a90-46f7-9554-21d30ce4c32d"),
		// 					PolicyDefinitionReferenceID: to.Ptr(""),
		// 					PolicySetDefinitionID: to.Ptr(""),
		// 					PolicySetDefinitionName: to.Ptr(""),
		// 					ResourceGroup: to.Ptr("myrg1"),
		// 					ResourceID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourceGroups/myrg1/providers/microsoft.insights/autoscalesettings/mytest1"),
		// 					ResourceLocation: to.Ptr("westus"),
		// 					ResourceTags: to.Ptr("tbd"),
		// 					ResourceType: to.Ptr("/microsoft.insights/autoscalesettings"),
		// 					SubscriptionID: to.Ptr("fffedd8f-ffff-fffd-fffd-fffed2f84852"),
		// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-10-09T17:41:47.000Z"); return t}()),
		// 			}},
		// 		}
	}
}
