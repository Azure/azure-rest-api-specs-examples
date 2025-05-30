package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/219b2e3ef270f18149774eb2793b48baacde982f/specification/resources/resource-manager/Microsoft.Authorization/preview/2022-07-01-preview/examples/createOrUpdatePolicyExemptionWithResourceSelectors.json
func ExampleExemptionsClient_CreateOrUpdate_createOrUpdateAPolicyExemptionWithResourceSelectors() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExemptionsClient().CreateOrUpdate(ctx, "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/resourceGroups/demoCluster", "DemoExpensiveVM", armpolicy.Exemption{
		Properties: &armpolicy.ExemptionProperties{
			Description:               to.Ptr("Exempt demo cluster from limit sku"),
			AssignmentScopeValidation: to.Ptr(armpolicy.AssignmentScopeValidationDefault),
			DisplayName:               to.Ptr("Exempt demo cluster"),
			ExemptionCategory:         to.Ptr(armpolicy.ExemptionCategoryWaiver),
			Metadata: map[string]any{
				"reason": "Temporary exemption for a expensive VM demo",
			},
			PolicyAssignmentID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyAssignments/CostManagement"),
			PolicyDefinitionReferenceIDs: []*string{
				to.Ptr("Limit_Skus")},
			ResourceSelectors: []*armpolicy.ResourceSelector{
				{
					Name: to.Ptr("SDPRegions"),
					Selectors: []*armpolicy.Selector{
						{
							In: []*string{
								to.Ptr("eastus2euap"),
								to.Ptr("centraluseuap")},
							Kind: to.Ptr(armpolicy.SelectorKindResourceLocation),
						}},
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Exemption = armpolicy.Exemption{
	// 	Name: to.Ptr("DemoExpensiveVM"),
	// 	Type: to.Ptr("Microsoft.Authorization/policyExemptions"),
	// 	ID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/resourceGroups/demoCluster/providers/Microsoft.Authorization/policyExemptions/DemoExpensiveVM"),
	// 	Properties: &armpolicy.ExemptionProperties{
	// 		Description: to.Ptr("Exempt demo cluster from limit sku"),
	// 		AssignmentScopeValidation: to.Ptr(armpolicy.AssignmentScopeValidationDefault),
	// 		DisplayName: to.Ptr("Exempt demo cluster"),
	// 		ExemptionCategory: to.Ptr(armpolicy.ExemptionCategoryWaiver),
	// 		Metadata: map[string]any{
	// 			"reason": "Temporary exemption for a expensive VM demo",
	// 		},
	// 		PolicyAssignmentID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyAssignments/CostManagement"),
	// 		PolicyDefinitionReferenceIDs: []*string{
	// 			to.Ptr("Limit_Skus")},
	// 			ResourceSelectors: []*armpolicy.ResourceSelector{
	// 				{
	// 					Name: to.Ptr("SDPRegions"),
	// 					Selectors: []*armpolicy.Selector{
	// 						{
	// 							In: []*string{
	// 								to.Ptr("eastus2euap"),
	// 								to.Ptr("centraluseuap")},
	// 								Kind: to.Ptr(armpolicy.SelectorKindResourceLocation),
	// 						}},
	// 				}},
	// 			},
	// 			SystemData: &armpolicy.SystemData{
	// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-01T01:01:01.107Z"); return t}()),
	// 				CreatedBy: to.Ptr("string"),
	// 				CreatedByType: to.Ptr(armpolicy.CreatedByTypeUser),
	// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-01T02:01:01.107Z"); return t}()),
	// 				LastModifiedBy: to.Ptr("string"),
	// 				LastModifiedByType: to.Ptr(armpolicy.CreatedByTypeUser),
	// 			},
	// 		}
}
