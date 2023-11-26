package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/05a9cdab363b8ec824094ee73950c04594325172/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyEvents_QueryResourceGroupScope.json
func ExamplePolicyEventsClient_NewListQueryResultsForResourceGroupPager_queryAtResourceGroupScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPolicyEventsClient().NewListQueryResultsForResourceGroupPager(armpolicyinsights.PolicyEventsResourceTypeDefault, "fffedd8f-ffff-fffd-fffd-fffed2f84852", "myResourceGroup", &armpolicyinsights.QueryOptions{Top: nil,
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
		// 	ODataContext: to.Ptr("https://management.azure.com/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourcegroups/myResourceGroup/providers/Microsoft.PolicyInsights/policyEvents/$metadata#default"),
		// 	ODataCount: to.Ptr[int32](2),
		// 	Value: []*armpolicyinsights.PolicyEvent{
		// 		{
		// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourcegroups/myResourceGroup/providers/Microsoft.PolicyInsights/policyEvents/$metadata#default/$entity"),
		// 			ComplianceState: to.Ptr("NonCompliant"),
		// 			IsCompliant: to.Ptr(false),
		// 			ManagementGroupIDs: to.Ptr("myManagementGroup,fff988bf-fff1-ffff-fffb-fffcd011db47"),
		// 			PolicyAssignmentID: to.Ptr("/providers/Microsoft.Management/managementGroups/myManagementGroup/providers/Microsoft.Authorization/policyAssignments/ec62f9b2a454487296f2ccd4"),
		// 			PolicyAssignmentName: to.Ptr("ec62f9b2a454487296f2ccd4"),
		// 			PolicyAssignmentOwner: to.Ptr("tbd"),
		// 			PolicyAssignmentParameters: to.Ptr("{\"ALLOWEDRESOURCEGROUPS_1\":{\"value\":[\"rg1\",\"rg2\"]},\"ALLOWEDRESOURCEGROUPS_2\":{\"value\":[\"myrg3\",\"myrg4\"]}}"),
		// 			PolicyAssignmentScope: to.Ptr("/providers/Microsoft.Management/managementGroups/myManagementGroup"),
		// 			PolicyDefinitionAction: to.Ptr("audit"),
		// 			PolicyDefinitionCategory: to.Ptr("tbd"),
		// 			PolicyDefinitionID: to.Ptr("/providers/Microsoft.Management/managementGroups/myManagementGroup/providers/Microsoft.Authorization/policyDefinitions/72c0c41a-c752-4bc0-9c61-0d6adc567066"),
		// 			PolicyDefinitionName: to.Ptr("72c0c41a-c752-4bc0-9c61-0d6adc567066"),
		// 			PolicyDefinitionReferenceID: to.Ptr("181565554491747128"),
		// 			PolicySetDefinitionID: to.Ptr("/providers/Microsoft.Management/managementGroups/myManagementGroup/providers/Microsoft.Authorization/policySetDefinitions/00b36c66-612b-44e2-9f8e-b758296d40fe"),
		// 			PolicySetDefinitionName: to.Ptr("00b36c66-612b-44e2-9f8e-b758296d40fe"),
		// 			PrincipalOid: to.Ptr("fffdfc0f-fff5-fff0-fff3-fff1a968dcc6"),
		// 			ResourceGroup: to.Ptr("myResourceGroup"),
		// 			ResourceID: to.Ptr("/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourcegroups/myResourceGroup/providers/Microsoft.ServiceFabric/clusters/myCluster/applications/myApplication"),
		// 			ResourceLocation: to.Ptr("eastus"),
		// 			ResourceTags: to.Ptr("tbd"),
		// 			ResourceType: to.Ptr("/Microsoft.ServiceFabric/clusters/applications"),
		// 			SubscriptionID: to.Ptr("fff10b27-fff3-fff5-fff8-fffbe01e86a5"),
		// 			TenantID: to.Ptr("fff988bf-fff1-ffff-fffb-fffcd011db47"),
		// 			Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-07T20:43:04.697Z"); return t}()),
		// 		},
		// 		{
		// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourcegroups/myResourceGroup/providers/Microsoft.PolicyInsights/policyEvents/$metadata#default/$entity"),
		// 			ComplianceState: to.Ptr("NonCompliant"),
		// 			IsCompliant: to.Ptr(false),
		// 			ManagementGroupIDs: to.Ptr("myManagementGroup,fff988bf-fff1-ffff-fffb-fffcd011db47"),
		// 			PolicyAssignmentID: to.Ptr("/providers/Microsoft.Management/managementGroups/myManagementGroup/providers/Microsoft.Authorization/policyAssignments/ec62f9b2a454487296f2ccd4"),
		// 			PolicyAssignmentName: to.Ptr("ec62f9b2a454487296f2ccd4"),
		// 			PolicyAssignmentOwner: to.Ptr("tbd"),
		// 			PolicyAssignmentParameters: to.Ptr("{\"ALLOWEDRESOURCEGROUPS_1\":{\"value\":[\"rg1\",\"rg2\"]},\"ALLOWEDRESOURCEGROUPS_2\":{\"value\":[\"myrg3\",\"myrg4\"]}}"),
		// 			PolicyAssignmentScope: to.Ptr("/providers/Microsoft.Management/managementGroups/myManagementGroup"),
		// 			PolicyDefinitionAction: to.Ptr("audit"),
		// 			PolicyDefinitionCategory: to.Ptr("tbd"),
		// 			PolicyDefinitionID: to.Ptr("/providers/Microsoft.Management/managementGroups/myManagementGroup/providers/Microsoft.Authorization/policyDefinitions/72c0c41a-c752-4bc0-9c61-0d6adc567066"),
		// 			PolicyDefinitionName: to.Ptr("72c0c41a-c752-4bc0-9c61-0d6adc567066"),
		// 			PolicyDefinitionReferenceID: to.Ptr("624540685646900425"),
		// 			PolicySetDefinitionID: to.Ptr("/providers/Microsoft.Management/managementGroups/myManagementGroup/providers/Microsoft.Authorization/policySetDefinitions/00b36c66-612b-44e2-9f8e-b758296d40fe"),
		// 			PolicySetDefinitionName: to.Ptr("00b36c66-612b-44e2-9f8e-b758296d40fe"),
		// 			PrincipalOid: to.Ptr("fffdfc0f-fff5-fff0-fff3-fff1a968dcc6"),
		// 			ResourceGroup: to.Ptr("myResourceGroup"),
		// 			ResourceID: to.Ptr("/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourcegroups/myResourceGroup/providers/Microsoft.ServiceFabric/clusters/myCluster/applications/myApplication"),
		// 			ResourceLocation: to.Ptr("eastus"),
		// 			ResourceTags: to.Ptr("tbd"),
		// 			ResourceType: to.Ptr("/Microsoft.ServiceFabric/clusters/applications"),
		// 			SubscriptionID: to.Ptr("fff10b27-fff3-fff5-fff8-fffbe01e86a5"),
		// 			TenantID: to.Ptr("fff988bf-fff1-ffff-fffb-fffcd011db47"),
		// 			Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-07T20:43:04.697Z"); return t}()),
		// 	}},
		// }
	}
}
