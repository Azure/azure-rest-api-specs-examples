package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c548b0bd279f5e233661b1c81fb5b61b19965cd/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/ComponentPolicyStates_QueryResourceGroupLevelPolicyAssignmentScope.json
func ExampleComponentPolicyStatesClient_ListQueryResultsForResourceGroupLevelPolicyAssignment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewComponentPolicyStatesClient().ListQueryResultsForResourceGroupLevelPolicyAssignment(ctx, "fffedd8f-ffff-fffd-fffd-fffed2f84852", "myResourceGroup", "myPolicyAssignment", armpolicyinsights.ComponentPolicyStatesResourceLatest, &armpolicyinsights.ComponentPolicyStatesClientListQueryResultsForResourceGroupLevelPolicyAssignmentOptions{Top: nil,
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
	// 	ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourceGroups/myResourceGroup/providers/Microsoft.Authorization/policyAssignments/myPolicyAssignment/providers/Microsoft.PolicyInsights/componentPolicyStates/$metadata#latest"),
	// 	ODataCount: to.Ptr[int32](2),
	// 	Value: []*armpolicyinsights.ComponentPolicyState{
	// 		{
	// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourceGroups/myResourceGroup/providers/Microsoft.Authorization/policyAssignments/myPolicyAssignment/providers/Microsoft.PolicyInsights/componentPolicyStates/$metadata#latest/$entity"),
	// 			ComplianceState: to.Ptr("NonCompliant"),
	// 			ComponentID: to.Ptr("cert-RSA-cert-3"),
	// 			ComponentName: to.Ptr("cert-RSA-cert-3"),
	// 			ComponentType: to.Ptr("Certificate"),
	// 			PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourceGroups/myResourceGroup/providers/Microsoft.Authorization/policyAssignments/myPolicyAssignment"),
	// 			PolicyAssignmentName: to.Ptr("myPolicyAssignment"),
	// 			PolicyAssignmentOwner: to.Ptr("tbd"),
	// 			PolicyAssignmentScope: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourceGroups/myResourceGroup"),
	// 			PolicyAssignmentVersion: to.Ptr("1.0.0"),
	// 			PolicyDefinitionAction: to.Ptr("audit"),
	// 			PolicyDefinitionCategory: to.Ptr("tbd"),
	// 			PolicyDefinitionGroupNames: []*string{
	// 				to.Ptr("myGroup")},
	// 				PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.Authorization/policyDefinitions/4a0425e4-97bf-4ad0-ab36-145b94083c60"),
	// 				PolicyDefinitionName: to.Ptr("4a0425e4-97bf-4ad0-ab36-145b94083c60"),
	// 				PolicyDefinitionVersion: to.Ptr("1.0.0-preview"),
	// 				ResourceGroup: to.Ptr("myResourceGroup"),
	// 				ResourceID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourceGroups/myResourceGroup/providers/Microsoft.Security/policies/mySecurityPolicy"),
	// 				ResourceType: to.Ptr("/Microsoft.KeyVault/vaults"),
	// 				SubscriptionID: to.Ptr("fffedd8f-ffff-fffd-fffd-fffed2f84852"),
	// 				Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-09T16:04:31.000Z"); return t}()),
	// 			},
	// 			{
	// 				AdditionalProperties: map[string]any{
	// 					"isCompliant": true,
	// 					"resourceTags": "tbd",
	// 				},
	// 				ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourceGroups/myResourceGroup/providers/Microsoft.Authorization/policyAssignments/myPolicyAssignment/providers/Microsoft.PolicyInsights/componentPolicyStates/$metadata#latest/$entity"),
	// 				ComplianceState: to.Ptr("Compliant"),
	// 				ComponentID: to.Ptr("cert-RSA-cert-3"),
	// 				ComponentName: to.Ptr("cert-RSA-cert-3"),
	// 				ComponentType: to.Ptr("Certificate"),
	// 				PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourceGroups/myResourceGroup/providers/Microsoft.Authorization/policyAssignments/myPolicyAssignment"),
	// 				PolicyAssignmentName: to.Ptr("myPolicyAssignment"),
	// 				PolicyAssignmentOwner: to.Ptr("tbd"),
	// 				PolicyAssignmentScope: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourceGroups/myResourceGroup"),
	// 				PolicyAssignmentVersion: to.Ptr("1.0.0"),
	// 				PolicyDefinitionAction: to.Ptr("audit"),
	// 				PolicyDefinitionCategory: to.Ptr("tbd"),
	// 				PolicyDefinitionGroupNames: []*string{
	// 					to.Ptr("myGroup")},
	// 					PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.Authorization/policyDefinitions/4a0425e4-97bf-4ad0-ab36-145b94083c60"),
	// 					PolicyDefinitionName: to.Ptr("4a0425e4-97bf-4ad0-ab36-145b94083c60"),
	// 					PolicyDefinitionVersion: to.Ptr("1.0.0-preview"),
	// 					ResourceGroup: to.Ptr("myResourceGroup"),
	// 					ResourceID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourceGroups/myResourceGroup/providers/Microsoft.Authorization/policyAssignments/test"),
	// 					ResourceType: to.Ptr("/Microsoft.KeyVault/vaults"),
	// 					SubscriptionID: to.Ptr("fffedd8f-ffff-fffd-fffd-fffed2f84852"),
	// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-09T16:04:31.000Z"); return t}()),
	// 			}},
	// 		}
}
