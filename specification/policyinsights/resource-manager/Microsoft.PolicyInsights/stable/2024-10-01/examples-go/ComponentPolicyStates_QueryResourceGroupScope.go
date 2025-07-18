package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c548b0bd279f5e233661b1c81fb5b61b19965cd/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/ComponentPolicyStates_QueryResourceGroupScope.json
func ExampleComponentPolicyStatesClient_ListQueryResultsForResourceGroup_queryLatestComponentPolicyStatesAtResourceGroupScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewComponentPolicyStatesClient().ListQueryResultsForResourceGroup(ctx, "fffedd8f-ffff-fffd-fffd-fffed2f84852", "myResourceGroup", armpolicyinsights.ComponentPolicyStatesResourceLatest, &armpolicyinsights.ComponentPolicyStatesClientListQueryResultsForResourceGroupOptions{Top: nil,
		OrderBy: nil,
		Select:  nil,
		From:    nil,
		To:      nil,
		Filter:  nil,
		Apply:   nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ComponentPolicyStatesQueryResults = armpolicyinsights.ComponentPolicyStatesQueryResults{
	// 	ODataContext: to.Ptr("https://management.azure.com/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourcegroups/myResourceGroup/providers/Microsoft.PolicyInsights/componentPolicyStates/$metadata#latest"),
	// 	ODataCount: to.Ptr[int32](2),
	// 	Value: []*armpolicyinsights.ComponentPolicyState{
	// 		{
	// 			AdditionalProperties: map[string]any{
	// 				"complianceReasonCode": "tbd",
	// 			},
	// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourcegroups/myResourceGroup/providers/Microsoft.PolicyInsights/componentPolicyStates/$metadata#latest/$entity"),
	// 			ComplianceState: to.Ptr("NonCompliant"),
	// 			ComponentID: to.Ptr("cert-RSA-cert-3"),
	// 			ComponentName: to.Ptr("cert-RSA-cert-3"),
	// 			ComponentType: to.Ptr("Certificate"),
	// 			PolicyAssignmentID: to.Ptr("/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/providers/Microsoft.Authorization/policyAssignments/test"),
	// 			PolicyAssignmentName: to.Ptr("test"),
	// 			PolicyAssignmentOwner: to.Ptr("tbd"),
	// 			PolicyAssignmentScope: to.Ptr("/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5"),
	// 			PolicyAssignmentVersion: to.Ptr("1.0.0"),
	// 			PolicyDefinitionAction: to.Ptr("audit"),
	// 			PolicyDefinitionCategory: to.Ptr("tbd"),
	// 			PolicyDefinitionGroupNames: []*string{
	// 				to.Ptr("myGroup")},
	// 				PolicyDefinitionID: to.Ptr("/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/providers/Microsoft.Authorization/policyDefinitions/ab108bc4-32df-4677-8b38-fa8b2905df56"),
	// 				PolicyDefinitionName: to.Ptr("ab108bc4-32df-4677-8b38-fa8b2905df56"),
	// 				PolicyDefinitionVersion: to.Ptr("1.0.0-preview"),
	// 				ResourceGroup: to.Ptr("myResourceGroup"),
	// 				ResourceID: to.Ptr("/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourcegroups/myResourceGroup/providers/Microsoft.KeyVault/Vaults/myKVName"),
	// 				ResourceLocation: to.Ptr("eastus"),
	// 				ResourceType: to.Ptr("/Microsoft.KeyVault/vaults"),
	// 				SubscriptionID: to.Ptr("fff10b27-fff3-fff5-fff8-fffbe01e86a5"),
	// 				Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-09T16:04:31.000Z"); return t}()),
	// 			},
	// 			{
	// 				AdditionalProperties: map[string]any{
	// 					"complianceReasonCode": "tbd",
	// 				},
	// 				ODataContext: to.Ptr("https://management.azure.com/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourcegroups/myResourceGroup/providers/Microsoft.PolicyInsights/componentPolicyStates/$metadata#latest/$entity"),
	// 				ComplianceState: to.Ptr("NonCompliant"),
	// 				ComponentID: to.Ptr("cert-RSA-cert-2"),
	// 				ComponentName: to.Ptr("cert-RSA-cert-2"),
	// 				ComponentType: to.Ptr("Certificate"),
	// 				PolicyAssignmentID: to.Ptr("/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/providers/Microsoft.Authorization/policyAssignments/test"),
	// 				PolicyAssignmentName: to.Ptr("test"),
	// 				PolicyAssignmentOwner: to.Ptr("tbd"),
	// 				PolicyAssignmentScope: to.Ptr("/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5"),
	// 				PolicyAssignmentVersion: to.Ptr("1.0.0"),
	// 				PolicyDefinitionAction: to.Ptr("audit"),
	// 				PolicyDefinitionCategory: to.Ptr("tbd"),
	// 				PolicyDefinitionGroupNames: []*string{
	// 					to.Ptr("myGroup")},
	// 					PolicyDefinitionID: to.Ptr("/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/providers/Microsoft.Authorization/policyDefinitions/ab108bc4-32df-4677-8b38-fa8b2905df59"),
	// 					PolicyDefinitionName: to.Ptr("ab108bc4-32df-4677-8b38-fa8b2905df59"),
	// 					PolicyDefinitionVersion: to.Ptr("1.0.0-preview"),
	// 					ResourceGroup: to.Ptr("myResourceGroup"),
	// 					ResourceID: to.Ptr("/subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourcegroups/myResourceGroup/providers/Microsoft.KeyVault/Vaults/myKVName"),
	// 					ResourceLocation: to.Ptr("eastus"),
	// 					ResourceType: to.Ptr("/Microsoft.KeyVault/vaults"),
	// 					SubscriptionID: to.Ptr("fff10b27-fff3-fff5-fff8-fffbe01e86a5"),
	// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-09T16:04:31.000Z"); return t}()),
	// 			}},
	// 		}
}
