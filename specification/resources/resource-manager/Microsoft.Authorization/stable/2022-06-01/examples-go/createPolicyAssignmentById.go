package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/resources/resource-manager/Microsoft.Authorization/stable/2022-06-01/examples/createPolicyAssignmentById.json
func ExampleAssignmentsClient_CreateByID_createOrUpdatePolicyAssignmentById() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAssignmentsClient().CreateByID(ctx, "providers/Microsoft.Management/managementGroups/MyManagementGroup/providers/Microsoft.Authorization/policyAssignments/LowCostStorage", armpolicy.Assignment{
		Properties: &armpolicy.AssignmentProperties{
			Description:     to.Ptr("Allow only storage accounts of SKU Standard_GRS or Standard_LRS to be created"),
			DisplayName:     to.Ptr("Enforce storage account SKU"),
			EnforcementMode: to.Ptr(armpolicy.EnforcementModeDefault),
			Metadata: map[string]any{
				"assignedBy": "Cheapskate Boss",
			},
			Parameters: map[string]*armpolicy.ParameterValuesValue{
				"listOfAllowedSKUs": {
					Value: []any{
						"Standard_GRS",
						"Standard_LRS",
					},
				},
			},
			PolicyDefinitionID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/7433c107-6db4-4ad1-b57a-a76dce0154a1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
