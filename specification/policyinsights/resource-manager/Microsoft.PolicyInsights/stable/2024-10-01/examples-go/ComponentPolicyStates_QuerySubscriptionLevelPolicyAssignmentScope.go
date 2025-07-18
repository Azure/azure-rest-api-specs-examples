package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c548b0bd279f5e233661b1c81fb5b61b19965cd/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/ComponentPolicyStates_QuerySubscriptionLevelPolicyAssignmentScope.json
func ExampleComponentPolicyStatesClient_ListQueryResultsForSubscriptionLevelPolicyAssignment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewComponentPolicyStatesClient().ListQueryResultsForSubscriptionLevelPolicyAssignment(ctx, "fffedd8f-ffff-fffd-fffd-fffed2f84852", "ec8f9645-8ecb-4abb-9c0b-5292f19d4003", armpolicyinsights.ComponentPolicyStatesResourceLatest, &armpolicyinsights.ComponentPolicyStatesClientListQueryResultsForSubscriptionLevelPolicyAssignmentOptions{Top: nil,
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
	// 	ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.Authorization/policyAssignments/ec8f9645-8ecb-4abb-9c0b-5292f19d4003/providers/Microsoft.PolicyInsights/policyStates/$metadata#latest"),
	// 	ODataCount: to.Ptr[int32](2),
	// 	Value: []*armpolicyinsights.ComponentPolicyState{
	// 		{
	// 			ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.Authorization/policyAssignments/ec8f9645-8ecb-4abb-9c0b-5292f19d4003/providers/Microsoft.PolicyInsights/policyStates/$metadata#latest/$entity"),
	// 			ComplianceState: to.Ptr("NonCompliant"),
	// 			ComponentID: to.Ptr("cert-RSA-cert-3"),
	// 			ComponentName: to.Ptr("cert-RSA-cert-3"),
	// 			ComponentType: to.Ptr("Certificate"),
	// 			PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.Authorization/policyAssignments/ec8f9645-8ecb-4abb-9c0b-5292f19d4003"),
	// 			PolicyAssignmentName: to.Ptr("ec8f9645-8ecb-4abb-9c0b-5292f19d4003"),
	// 			PolicyAssignmentOwner: to.Ptr("tbd"),
	// 			PolicyAssignmentScope: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852"),
	// 			PolicyAssignmentVersion: to.Ptr("1.0.0"),
	// 			PolicyDefinitionAction: to.Ptr("audit"),
	// 			PolicyDefinitionCategory: to.Ptr("tbd"),
	// 			PolicyDefinitionGroupNames: []*string{
	// 				to.Ptr("myGroup")},
	// 				PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.Authorization/policyDefinitions/c8b79b49-a579-4045-984e-1b249ab8b474"),
	// 				PolicyDefinitionName: to.Ptr("c8b79b49-a579-4045-984e-1b249ab8b474"),
	// 				PolicyDefinitionReferenceID: to.Ptr("2124621540977569058"),
	// 				PolicyDefinitionVersion: to.Ptr("1.0.0-preview"),
	// 				PolicySetDefinitionVersion: to.Ptr("2.0.1"),
	// 				ResourceGroup: to.Ptr("myResourceGroup"),
	// 				ResourceID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVault"),
	// 				ResourceLocation: to.Ptr("eastus"),
	// 				ResourceType: to.Ptr("/Microsoft.KeyVault/vaults"),
	// 				SubscriptionID: to.Ptr("fffedd8f-ffff-fffd-fffd-fffed2f84852"),
	// 				Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-09T16:04:31.000Z"); return t}()),
	// 			},
	// 			{
	// 				AdditionalProperties: map[string]any{
	// 					"isCompliant": true,
	// 				},
	// 				ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.Authorization/policyAssignments/ec8f9645-8ecb-4abb-9c0b-5292f19d4003/providers/Microsoft.PolicyInsights/policyStates/$metadata#latest/$entity"),
	// 				ComplianceState: to.Ptr("Compliant"),
	// 				ComponentID: to.Ptr("cert-RSA-cert-3"),
	// 				ComponentName: to.Ptr("cert-RSA-cert-3"),
	// 				ComponentType: to.Ptr("Certificate"),
	// 				PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.Authorization/policyAssignments/ec8f9645-8ecb-4abb-9c0b-5292f19d4003"),
	// 				PolicyAssignmentName: to.Ptr("ec8f9645-8ecb-4abb-9c0b-5292f19d4003"),
	// 				PolicyAssignmentOwner: to.Ptr("tbd"),
	// 				PolicyAssignmentScope: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852"),
	// 				PolicyAssignmentVersion: to.Ptr("1.0.0"),
	// 				PolicyDefinitionAction: to.Ptr("audit"),
	// 				PolicyDefinitionCategory: to.Ptr("tbd"),
	// 				PolicyDefinitionGroupNames: []*string{
	// 					to.Ptr("myGroup")},
	// 					PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.Authorization/policyDefinitions/24813039-7534-408a-9842-eb99f45721b1"),
	// 					PolicyDefinitionName: to.Ptr("24813039-7534-408a-9842-eb99f45721b1"),
	// 					PolicyDefinitionReferenceID: to.Ptr("14799174781370023846"),
	// 					PolicyDefinitionVersion: to.Ptr("1.0.0-preview"),
	// 					PolicySetDefinitionVersion: to.Ptr("2.0.1"),
	// 					ResourceGroup: to.Ptr("myResourceGroup"),
	// 					ResourceID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVault"),
	// 					ResourceLocation: to.Ptr("eastus"),
	// 					ResourceType: to.Ptr("/Microsoft.KeyVault/vaults"),
	// 					SubscriptionID: to.Ptr("fffedd8f-ffff-fffd-fffd-fffed2f84852"),
	// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-09T16:04:31.000Z"); return t}()),
	// 			}},
	// 		}
}
