package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/05a9cdab363b8ec824094ee73950c04594325172/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_QueryNestedResourceScope.json
func ExamplePolicyStatesClient_NewListQueryResultsForResourcePager_queryAllPolicyStatesAtNestedResourceScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPolicyStatesClient().NewListQueryResultsForResourcePager(armpolicyinsights.PolicyStatesResourceDefault, "subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourceGroups/myResourceGroup/providers/Microsoft.ServiceFabric/clusters/myCluster/applications/myApplication", &armpolicyinsights.QueryOptions{Top: nil,
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
		// page.PolicyStatesQueryResults = armpolicyinsights.PolicyStatesQueryResults{
		// 	ODataContext: to.Ptr("https://management.azure.com/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourcegroups/myResourceGroup/providers/Microsoft.ServiceFabric/clusters/myCluster/applications/myApplication/providers/Microsoft.PolicyInsights/policyStates/$metadata#default"),
		// 	ODataCount: to.Ptr[int32](2),
		// 	Value: []*armpolicyinsights.PolicyState{
		// 		{
		// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourcegroups/myResourceGroup/providers/Microsoft.ServiceFabric/clusters/myCluster/applications/myApplication/providers/Microsoft.PolicyInsights/policyStates/$metadata#default/$entity"),
		// 			ComplianceState: to.Ptr("NonCompliant"),
		// 			IsCompliant: to.Ptr(false),
		// 			ManagementGroupIDs: to.Ptr("mymg,fff988bf-fff1-ffff-fffb-fffcd011db47"),
		// 			PolicyAssignmentID: to.Ptr("/providers/Microsoft.Management/managementGroups/mymg/providers/Microsoft.Authorization/policyAssignments/186044306c044a1d8c0ff76c"),
		// 			PolicyAssignmentName: to.Ptr("186044306c044a1d8c0ff76c"),
		// 			PolicyAssignmentOwner: to.Ptr("tbd"),
		// 			PolicyAssignmentParameters: to.Ptr("{\"allowedLocations\":{\"value\":[\"centralus\"]}}"),
		// 			PolicyAssignmentScope: to.Ptr("/providers/Microsoft.Management/managementGroups/mymg"),
		// 			PolicyAssignmentVersion: to.Ptr("1.0.0"),
		// 			PolicyDefinitionAction: to.Ptr("audit"),
		// 			PolicyDefinitionCategory: to.Ptr("tbd"),
		// 			PolicyDefinitionGroupNames: []*string{
		// 				to.Ptr("myGroup")},
		// 				PolicyDefinitionID: to.Ptr("/providers/Microsoft.Management/managementGroups/mymg/providers/Microsoft.Authorization/policyDefinitions/022d9357-5a90-46f7-9554-21d30ce4c32d"),
		// 				PolicyDefinitionName: to.Ptr("022d9357-5a90-46f7-9554-21d30ce4c32d"),
		// 				PolicyDefinitionVersion: to.Ptr("1.0.0-preview"),
		// 				ResourceGroup: to.Ptr("myResourceGroup"),
		// 				ResourceID: to.Ptr("/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourcegroups/myResourceGroup/providers/Microsoft.ServiceFabric/clusters/myCluster/applications/myApplication"),
		// 				ResourceLocation: to.Ptr("eastus"),
		// 				ResourceTags: to.Ptr("tbd"),
		// 				ResourceType: to.Ptr("/Microsoft.ServiceFabric/clusters/applications"),
		// 				SubscriptionID: to.Ptr("fff10b27-fff3-fff5-fff8-fffbe01e86a5"),
		// 				Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-10-09T16:04:31Z"); return t}()),
		// 			},
		// 			{
		// 				ODataContext: to.Ptr("https://management.azure.com/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourcegroups/myResourceGroup/providers/Microsoft.ServiceFabric/clusters/myCluster/applications/myApplication/providers/Microsoft.PolicyInsights/policyStates/$metadata#default/$entity"),
		// 				ComplianceState: to.Ptr("Compliant"),
		// 				IsCompliant: to.Ptr(true),
		// 				ManagementGroupIDs: to.Ptr("mymg,fff988bf-fff1-ffff-fffb-fffcd011db47"),
		// 				PolicyAssignmentID: to.Ptr("/providers/Microsoft.Management/managementGroups/mymg/providers/Microsoft.Authorization/policyAssignments/186044306c044a1d8c0ff76c"),
		// 				PolicyAssignmentName: to.Ptr("186044306c044a1d8c0ff76c"),
		// 				PolicyAssignmentOwner: to.Ptr("tbd"),
		// 				PolicyAssignmentParameters: to.Ptr("{\"allowedLocations\":{\"value\":[\"centralus\"]}}"),
		// 				PolicyAssignmentScope: to.Ptr("/providers/Microsoft.Management/managementGroups/mymg"),
		// 				PolicyAssignmentVersion: to.Ptr("1.0.0"),
		// 				PolicyDefinitionAction: to.Ptr("audit"),
		// 				PolicyDefinitionCategory: to.Ptr("tbd"),
		// 				PolicyDefinitionGroupNames: []*string{
		// 					to.Ptr("myGroup")},
		// 					PolicyDefinitionID: to.Ptr("/providers/Microsoft.Management/managementGroups/mymg/providers/Microsoft.Authorization/policyDefinitions/022d9357-5a90-46f7-9554-21d30ce4c32d"),
		// 					PolicyDefinitionName: to.Ptr("022d9357-5a90-46f7-9554-21d30ce4c32d"),
		// 					PolicyDefinitionVersion: to.Ptr("1.0.0-preview"),
		// 					ResourceGroup: to.Ptr("myResourceGroup"),
		// 					ResourceID: to.Ptr("/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourcegroups/myResourceGroup/providers/Microsoft.ServiceFabric/clusters/myCluster/applications/myApplication"),
		// 					ResourceLocation: to.Ptr("eastus"),
		// 					ResourceTags: to.Ptr("tbd"),
		// 					ResourceType: to.Ptr("/Microsoft.ServiceFabric/clusters/applications"),
		// 					SubscriptionID: to.Ptr("fff10b27-fff3-fff5-fff8-fffbe01e86a5"),
		// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-10-09T16:04:31Z"); return t}()),
		// 			}},
		// 		}
	}
}
