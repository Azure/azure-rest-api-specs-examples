package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/219b2e3ef270f18149774eb2793b48baacde982f/specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/createPolicyAssignmentWithOverrides.json
func ExampleAssignmentsClient_Create_createOrUpdateAPolicyAssignmentWithOverrides() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAssignmentsClient().Create(ctx, "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2", "CostManagement", armpolicy.Assignment{
		Properties: &armpolicy.AssignmentProperties{
			Description: to.Ptr("Limit the resource location and resource SKU"),
			DisplayName: to.Ptr("Limit the resource location and resource SKU"),
			Metadata: map[string]any{
				"assignedBy": "Special Someone",
			},
			Overrides: []*armpolicy.Override{
				{
					Kind: to.Ptr(armpolicy.OverrideKindPolicyEffect),
					Selectors: []*armpolicy.Selector{
						{
							In: []*string{
								to.Ptr("Limit_Skus"),
								to.Ptr("Limit_Locations")},
							Kind: to.Ptr(armpolicy.SelectorKindPolicyDefinitionReferenceID),
						}},
					Value: to.Ptr("Audit"),
				}},
			PolicyDefinitionID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policySetDefinitions/CostManagement"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
