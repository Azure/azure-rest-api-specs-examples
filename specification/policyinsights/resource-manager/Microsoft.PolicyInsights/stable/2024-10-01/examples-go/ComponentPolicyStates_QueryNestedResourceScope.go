package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c548b0bd279f5e233661b1c81fb5b61b19965cd/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/ComponentPolicyStates_QueryNestedResourceScope.json
func ExampleComponentPolicyStatesClient_ListQueryResultsForResource_queryLatestComponentPolicyStatesAtNestedResourceScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewComponentPolicyStatesClient().ListQueryResultsForResource(ctx, "subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVault", armpolicyinsights.ComponentPolicyStatesResourceLatest, &armpolicyinsights.ComponentPolicyStatesClientListQueryResultsForResourceOptions{Top: nil,
		OrderBy: nil,
		Select:  nil,
		From:    nil,
		To:      nil,
		Filter:  nil,
		Apply:   nil,
		Expand:  nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ComponentPolicyStatesQueryResults = armpolicyinsights.ComponentPolicyStatesQueryResults{
	// 	ODataContext: to.Ptr("https://management.azure.com/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVault/providers/Microsoft.PolicyInsights/policyStates/$metadata#latest"),
	// 	ODataCount: to.Ptr[int32](2),
	// 	Value: []*armpolicyinsights.ComponentPolicyState{
	// 		{
	// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVault/providers/Microsoft.PolicyInsights/policyStates/$metadata#latest/$entity"),
	// 			ComplianceState: to.Ptr("NonCompliant"),
	// 			ComponentID: to.Ptr("cert-RSA-cert-3"),
	// 			ComponentName: to.Ptr("cert-RSA-cert-3"),
	// 			ComponentType: to.Ptr("Certificate"),
	// 			PolicyAssignmentID: to.Ptr("/providers/Microsoft.Management/managementGroups/mymg/providers/Microsoft.Authorization/policyAssignments/186044306c044a1d8c0ff76c"),
	// 			PolicyAssignmentName: to.Ptr("186044306c044a1d8c0ff76c"),
	// 			PolicyAssignmentOwner: to.Ptr("tbd"),
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
	// 				ResourceID: to.Ptr("subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVault"),
	// 				ResourceLocation: to.Ptr("eastus"),
	// 				ResourceType: to.Ptr("/Microsoft.KeyVault/vaults"),
	// 				SubscriptionID: to.Ptr("fff10b27-fff3-fff5-fff8-fffbe01e86a5"),
	// 				Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-09T16:04:31.000Z"); return t}()),
	// 			},
	// 			{
	// 				AdditionalProperties: map[string]any{
	// 					"isCompliant": true,
	// 				},
	// 				ODataContext: to.Ptr("https://management.azure.com/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVault/providers/Microsoft.PolicyInsights/policyStates/$metadata#latest/$entity"),
	// 				ComplianceState: to.Ptr("Compliant"),
	// 				ComponentID: to.Ptr("cert-RSA-cert-3"),
	// 				ComponentName: to.Ptr("cert-RSA-cert-3"),
	// 				ComponentType: to.Ptr("Certificate"),
	// 				PolicyAssignmentID: to.Ptr("/providers/Microsoft.Management/managementGroups/mymg/providers/Microsoft.Authorization/policyAssignments/186044306c044a1d8c0ff76c"),
	// 				PolicyAssignmentName: to.Ptr("186044306c044a1d8c0ff76c"),
	// 				PolicyAssignmentOwner: to.Ptr("tbd"),
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
	// 					ResourceID: to.Ptr("/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVault"),
	// 					ResourceLocation: to.Ptr("eastus"),
	// 					ResourceType: to.Ptr("/Microsoft.KeyVault/vaults/myVault"),
	// 					SubscriptionID: to.Ptr("fff10b27-fff3-fff5-fff8-fffbe01e86a5"),
	// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-09T16:04:31.000Z"); return t}()),
	// 			}},
	// 		}
}
