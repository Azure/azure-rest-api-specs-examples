package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy/v2"
)

// Generated from example definition: 2026-01-01-preview/getPolicyExemptionWithResourceSelectors.json
func ExampleExemptionsClient_Get_retrieveAPolicyExemptionWithResourceSelectors() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExemptionsClient().Get(ctx, "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/resourceGroups/demoCluster", "DemoExpensiveVM", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpolicy.ExemptionsClientGetResponse{
	// 	Exemption: armpolicy.Exemption{
	// 		Properties: &armpolicy.ExemptionProperties{
	// 			PolicyAssignmentID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyAssignments/CostManagement"),
	// 			PolicyDefinitionReferenceIDs: []*string{
	// 				to.Ptr("Limit_Skus"),
	// 			},
	// 			ExemptionCategory: to.Ptr(armpolicy.ExemptionCategoryWaiver),
	// 			DisplayName: to.Ptr("Exempt demo cluster"),
	// 			Description: to.Ptr("Exempt demo cluster from limit sku"),
	// 			Metadata: map[string]any{
	// 				"reason": "Temporary exemption for a expensive VM demo",
	// 			},
	// 			AssignmentScopeValidation: to.Ptr(armpolicy.AssignmentScopeValidationDefault),
	// 			ResourceSelectors: []*armpolicy.ResourceSelector{
	// 				{
	// 					Name: to.Ptr("SDPRegions"),
	// 					Selectors: []*armpolicy.Selector{
	// 						{
	// 							Kind: to.Ptr(armpolicy.SelectorKindResourceLocation),
	// 							In: []*string{
	// 								to.Ptr("eastus2euap"),
	// 								to.Ptr("centraluseuap"),
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		SystemData: &armpolicy.SystemData{
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armpolicy.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-01T01:01:01.1075056Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armpolicy.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-01T02:01:01.1075056Z"); return t}()),
	// 		},
	// 		ID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/resourceGroups/demoCluster/providers/Microsoft.Authorization/policyExemptions/DemoExpensiveVM"),
	// 		Type: to.Ptr("Microsoft.Authorization/policyExemptions"),
	// 		Name: to.Ptr("DemoExpensiveVM"),
	// 	},
	// }
}
