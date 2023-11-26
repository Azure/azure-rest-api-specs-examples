package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/05a9cdab363b8ec824094ee73950c04594325172/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_QueryResourceScopeExpandComponents.json
func ExamplePolicyStatesClient_NewListQueryResultsForResourcePager_queryComponentPolicyComplianceStateAtResourceScopeFilteredByGivenAssignment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPolicyStatesClient().NewListQueryResultsForResourcePager(armpolicyinsights.PolicyStatesResourceLatest, "subscriptions/e78961ba-36fe-4739-9212-e3031b4c8db7/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/Vaults/myKVName", &armpolicyinsights.QueryOptions{Top: nil,
		Filter:    to.Ptr("policyAssignmentId eq '/subscriptions/e78961ba-36fe-4739-9212-e3031b4c8db7/providers/microsoft.authorization/policyassignments/560050f83dbb4a24974323f8'"),
		OrderBy:   nil,
		Select:    nil,
		From:      nil,
		To:        nil,
		Apply:     nil,
		SkipToken: nil,
		Expand:    to.Ptr("components($filter=ComplianceState eq 'NonCompliant' or ComplianceState eq 'Compliant')"),
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
		// 	ODataContext: to.Ptr("https://management.azure.com/subscriptions/e78961ba-36fe-4739-9212-e3031b4c8db7/resourcegroups/myResourceGroup/providers/Microsoft.KeyVault/Vaults/myKVName/providers/Microsoft.PolicyInsights/policyStates/$metadata#latest"),
		// 	ODataCount: to.Ptr[int32](1),
		// 	Value: []*armpolicyinsights.PolicyState{
		// 		{
		// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/e78961ba-36fe-4739-9212-e3031b4c8db7/resourcegroups/myResourceGroup/providers/Microsoft.KeyVault/Vaults/myKVName/providers/Microsoft.PolicyInsights/policyStates/$metadata#latest/$entity"),
		// 			ComplianceState: to.Ptr("NonCompliant"),
		// 			Components: []*armpolicyinsights.ComponentStateDetails{
		// 				{
		// 					Name: to.Ptr("cert-RSA-cert-3"),
		// 					Type: to.Ptr("Certificate"),
		// 					ComplianceState: to.Ptr("NonCompliant"),
		// 					ID: to.Ptr("cert-RSA-cert-3"),
		// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-21T17:45:23.036Z"); return t}()),
		// 				},
		// 				{
		// 					Name: to.Ptr("cert-RSA-cert-2"),
		// 					Type: to.Ptr("Certificate"),
		// 					ComplianceState: to.Ptr("NonCompliant"),
		// 					ID: to.Ptr("cert-RSA-cert-2"),
		// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-21T17:45:22.768Z"); return t}()),
		// 				},
		// 				{
		// 					Name: to.Ptr("cert-rsa-cert-expires-2months"),
		// 					Type: to.Ptr("Certificate"),
		// 					ComplianceState: to.Ptr("NonCompliant"),
		// 					ID: to.Ptr("cert-rsa-cert-expires-2months"),
		// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-21T17:45:22.684Z"); return t}()),
		// 				},
		// 				{
		// 					Name: to.Ptr("cert-RSA-cert-1"),
		// 					Type: to.Ptr("Certificate"),
		// 					ComplianceState: to.Ptr("Compliant"),
		// 					ID: to.Ptr("cert-RSA-cert-1"),
		// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-21T17:45:22.650Z"); return t}()),
		// 				},
		// 				{
		// 					Name: to.Ptr("cert-rsa-cert-expires-1month"),
		// 					Type: to.Ptr("Certificate"),
		// 					ComplianceState: to.Ptr("NonCompliant"),
		// 					ID: to.Ptr("cert-rsa-cert-expires-1month"),
		// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-21T17:45:22.552Z"); return t}()),
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
		// 			ResourceGroup: to.Ptr("myResourceGroup"),
		// 			ResourceID: to.Ptr("/subscriptions/e78961ba-36fe-4739-9212-e3031b4c8db7/resourcegroups/myResourceGroup/providers/microsoft.keyvault/vaults/myKVName"),
		// 			ResourceLocation: to.Ptr("westcentralus"),
		// 			ResourceTags: to.Ptr("tbd"),
		// 			ResourceType: to.Ptr("Microsoft.KeyVault/vaults"),
		// 			SubscriptionID: to.Ptr("e78961ba-36fe-4739-9212-e3031b4c8db7"),
		// 			Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-21T19:42:08.325Z"); return t}()),
		// 	}},
		// }
	}
}
