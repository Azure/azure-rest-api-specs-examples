package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy/v2"
)

// Generated from example definition: 2026-01-01-preview/createPolicyAssignmentWithSelfserveExemptionSettings.json
func ExampleAssignmentsClient_Create_createOrUpdateAPolicyAssignmentWithSelfServeExemptionSettings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAssignmentsClient().Create(ctx, "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2", "CostManagement", armpolicy.Assignment{
		Location: to.Ptr("eastus"),
		Properties: &armpolicy.AssignmentProperties{
			Description:       to.Ptr("Limit the resource location and resource SKU"),
			DefinitionVersion: to.Ptr("1.*.*"),
			DisplayName:       to.Ptr("Limit the resource location and resource SKU"),
			EnforcementMode:   to.Ptr(armpolicy.EnforcementModeDefault),
			Metadata: map[string]any{
				"assignedBy": "Foo Bar",
			},
			PolicyDefinitionID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policySetDefinitions/CostManagement"),
			SelfServeExemptionSettings: &armpolicy.SelfServeExemptionSettings{
				Enabled: to.Ptr(true),
				PolicyDefinitionReferenceIDs: []*string{
					to.Ptr("Limit_Skus"),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
