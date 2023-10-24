package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/resources/resource-manager/Microsoft.Authorization/stable/2022-06-01/examples/createPolicyAssignmentWithResourceSelectors.json
func ExampleAssignmentsClient_Create_createOrUpdateAPolicyAssignmentWithResourceSelectors() {
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
			PolicyDefinitionID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policySetDefinitions/CostManagement"),
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
}
