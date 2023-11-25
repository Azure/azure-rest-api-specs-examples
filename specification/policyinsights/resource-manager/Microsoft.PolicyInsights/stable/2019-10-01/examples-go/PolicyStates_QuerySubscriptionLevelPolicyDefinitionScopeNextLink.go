package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/05a9cdab363b8ec824094ee73950c04594325172/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_QuerySubscriptionLevelPolicyDefinitionScopeNextLink.json
func ExamplePolicyStatesClient_NewListQueryResultsForPolicyDefinitionPager_queryLatestAtSubscriptionLevelPolicyDefinitionScopeWithNextLink() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPolicyStatesClient().NewListQueryResultsForPolicyDefinitionPager(armpolicyinsights.PolicyStatesResourceLatest, "fffedd8f-ffff-fffd-fffd-fffed2f84852", "24813039-7534-408a-9842-eb99f45721b1", &armpolicyinsights.QueryOptions{Top: nil,
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
		// 	ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.Authorization/policyDefinitions/24813039-7534-408a-9842-eb99f45721b1/providers/Microsoft.PolicyInsights/policyStates/$metadata#latest"),
		// 	ODataCount: to.Ptr[int32](2),
		// 	Value: []*armpolicyinsights.PolicyState{
		// 		{
		// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.Authorization/policyDefinitions/24813039-7534-408a-9842-eb99f45721b1/providers/Microsoft.PolicyInsights/policyStates/$metadata#latest/$entity"),
		// 			ComplianceState: to.Ptr("NonCompliant"),
		// 			IsCompliant: to.Ptr(false),
		// 			ManagementGroupIDs: to.Ptr("mymg,fff988bf-fff1-ffff-fffb-fffcd011db47"),
		// 			PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.Authorization/policyAssignments/d9da7e80af6344ab9d342aa7"),
		// 			PolicyAssignmentName: to.Ptr("d9da7e80af6344ab9d342aa7"),
		// 			PolicyAssignmentOwner: to.Ptr("tbd"),
		// 			PolicyAssignmentParameters: to.Ptr("{\"tagName\":{\"value\":\"MyTag\"}}"),
		// 			PolicyAssignmentScope: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852"),
		// 			PolicyDefinitionAction: to.Ptr("audit"),
		// 			PolicyDefinitionCategory: to.Ptr("tbd"),
		// 			PolicyDefinitionGroupNames: []*string{
		// 				to.Ptr("myGroup")},
		// 				PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.Authorization/policyDefinitions/24813039-7534-408a-9842-eb99f45721b1"),
		// 				PolicyDefinitionName: to.Ptr("24813039-7534-408a-9842-eb99f45721b1"),
		// 				ResourceGroup: to.Ptr("myResourceGroup"),
		// 				ResourceID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourcegroups/myResourceGroup/providers/Microsoft.ServiceFabric/clusters/myCluster/applications/myApplication/services/myService"),
		// 				ResourceLocation: to.Ptr("eastus"),
		// 				ResourceTags: to.Ptr("tbd"),
		// 				ResourceType: to.Ptr("/Microsoft.ServiceFabric/clusters/applications/services"),
		// 				SubscriptionID: to.Ptr("fffedd8f-ffff-fffd-fffd-fffed2f84852"),
		// 				Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-10-13T00:50:27.000Z"); return t}()),
		// 			},
		// 			{
		// 				ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.Authorization/policyDefinitions/24813039-7534-408a-9842-eb99f45721b1/providers/Microsoft.PolicyInsights/policyStates/$metadata#latest/$entity"),
		// 				ComplianceState: to.Ptr("Compliant"),
		// 				IsCompliant: to.Ptr(true),
		// 				ManagementGroupIDs: to.Ptr("mymg,fff988bf-fff1-ffff-fffb-fffcd011db47"),
		// 				PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.Authorization/policyAssignments/d9da7e80af6344ab9d342aa7"),
		// 				PolicyAssignmentName: to.Ptr("d9da7e80af6344ab9d342aa7"),
		// 				PolicyAssignmentOwner: to.Ptr("tbd"),
		// 				PolicyAssignmentParameters: to.Ptr("{\"tagName\":{\"value\":\"MyTag\"}}"),
		// 				PolicyAssignmentScope: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852"),
		// 				PolicyDefinitionAction: to.Ptr("audit"),
		// 				PolicyDefinitionCategory: to.Ptr("tbd"),
		// 				PolicyDefinitionGroupNames: []*string{
		// 					to.Ptr("myGroup")},
		// 					PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.Authorization/policyDefinitions/24813039-7534-408a-9842-eb99f45721b1"),
		// 					PolicyDefinitionName: to.Ptr("24813039-7534-408a-9842-eb99f45721b1"),
		// 					ResourceGroup: to.Ptr("myResourceGroup"),
		// 					ResourceID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourcegroups/myResourceGroup/providers/Microsoft.ServiceFabric/clusters/myCluster/applications/myApplication/services/myService"),
		// 					ResourceLocation: to.Ptr("eastus"),
		// 					ResourceTags: to.Ptr("tbd"),
		// 					ResourceType: to.Ptr("/Microsoft.ServiceFabric/clusters/applications/services"),
		// 					SubscriptionID: to.Ptr("fffedd8f-ffff-fffd-fffd-fffed2f84852"),
		// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-10-13T00:50:27.000Z"); return t}()),
		// 			}},
		// 		}
	}
}
