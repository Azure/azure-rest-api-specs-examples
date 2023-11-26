package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/05a9cdab363b8ec824094ee73950c04594325172/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyEvents_QueryResourceGroupLevelPolicyAssignmentScopeNextLink.json
func ExamplePolicyEventsClient_NewListQueryResultsForResourceGroupLevelPolicyAssignmentPager_queryAtResourceGroupLevelPolicyAssignmentScopeWithNextLink() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPolicyEventsClient().NewListQueryResultsForResourceGroupLevelPolicyAssignmentPager(armpolicyinsights.PolicyEventsResourceTypeDefault, "fffedd8f-ffff-fffd-fffd-fffed2f84852", "myResourceGroup", "myPolicyAssignment", &armpolicyinsights.QueryOptions{Top: nil,
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
		// page.PolicyEventsQueryResults = armpolicyinsights.PolicyEventsQueryResults{
		// 	ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourcegroups/myResourceGroup/providers/Microsoft.Authorization/policyAssignments/myPolicyAssignment/providers/Microsoft.PolicyInsights/policyEvents/$metadata#default"),
		// 	ODataCount: to.Ptr[int32](2),
		// 	Value: []*armpolicyinsights.PolicyEvent{
		// 		{
		// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourcegroups/myResourceGroup/providers/Microsoft.Authorization/policyAssignments/myPolicyAssignment/providers/Microsoft.PolicyInsights/policyEvents/$metadata#default/$entity"),
		// 			ComplianceState: to.Ptr("NonCompliant"),
		// 			IsCompliant: to.Ptr(false),
		// 			ManagementGroupIDs: to.Ptr("mymg,fff988bf-fff1-ffff-fffb-fffcd011db47"),
		// 			PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourcegroups/myResourceGroup/providers/Microsoft.Authorization/policyAssignments/myPolicyAssignment"),
		// 			PolicyAssignmentName: to.Ptr("myPolicyAssignment"),
		// 			PolicyAssignmentOwner: to.Ptr("tbd"),
		// 			PolicyAssignmentScope: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourcegroups/myResourceGroup"),
		// 			PolicyDefinitionAction: to.Ptr("deny"),
		// 			PolicyDefinitionCategory: to.Ptr("tbd"),
		// 			PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyDefinitions/myPolicyDefinition"),
		// 			PolicyDefinitionName: to.Ptr("myPolicyAssignment"),
		// 			PrincipalOid: to.Ptr("fff2f355-fff2-fffc-fffb-fff1639dff94"),
		// 			ResourceGroup: to.Ptr("myResourceGroup"),
		// 			ResourceID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourcegroups/myResourceGroup/providers/microsoft.operationalinsights/workspaces/defaultworkspace-fffedd8f-ffff-fffd-fffd-fffed2f84852-eus"),
		// 			ResourceLocation: to.Ptr("eastus"),
		// 			ResourceTags: to.Ptr("tbd"),
		// 			ResourceType: to.Ptr("/microsoft.operationalinsights/workspaces"),
		// 			SubscriptionID: to.Ptr("fffedd8f-ffff-fffd-fffd-fffed2f84852"),
		// 			TenantID: to.Ptr("fff988bf-fff1-ffff-fffb-fffcd011db47"),
		// 			Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-08T00:07:16.280Z"); return t}()),
		// 		},
		// 		{
		// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourcegroups/myResourceGroup/providers/Microsoft.Authorization/policyAssignments/myPolicyAssignment/providers/Microsoft.PolicyInsights/policyEvents/$metadata#default/$entity"),
		// 			ComplianceState: to.Ptr("NonCompliant"),
		// 			IsCompliant: to.Ptr(false),
		// 			ManagementGroupIDs: to.Ptr("mymg,fff988bf-fff1-ffff-fffb-fffcd011db47"),
		// 			PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourcegroups/myResourceGroup/providers/Microsoft.Authorization/policyAssignments/myPolicyAssignment"),
		// 			PolicyAssignmentName: to.Ptr("myPolicyAssignment"),
		// 			PolicyAssignmentOwner: to.Ptr("tbd"),
		// 			PolicyAssignmentScope: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourcegroups/myResourceGroup"),
		// 			PolicyDefinitionAction: to.Ptr("deny"),
		// 			PolicyDefinitionCategory: to.Ptr("tbd"),
		// 			PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyDefinitions/myPolicyDefinition"),
		// 			PolicyDefinitionName: to.Ptr("myPolicyAssignment"),
		// 			PrincipalOid: to.Ptr("fff2f355-fff2-fffc-fffb-fff1639dff94"),
		// 			ResourceGroup: to.Ptr("myResourceGroup"),
		// 			ResourceID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourcegroups/myResourceGroup/providers/microsoft.operationalinsights/workspaces/defaultworkspace-fffedd8f-ffff-fffd-fffd-fffed2f84852-eus"),
		// 			ResourceLocation: to.Ptr("eastus"),
		// 			ResourceTags: to.Ptr("tbd"),
		// 			ResourceType: to.Ptr("/microsoft.operationalinsights/workspaces"),
		// 			SubscriptionID: to.Ptr("fffedd8f-ffff-fffd-fffd-fffed2f84852"),
		// 			TenantID: to.Ptr("fff988bf-fff1-ffff-fffb-fffcd011db47"),
		// 			Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-08T00:06:08.430Z"); return t}()),
		// 	}},
		// }
	}
}
