package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/219b2e3ef270f18149774eb2793b48baacde982f/specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/deletePolicyAssignmentById.json
func ExampleAssignmentsClient_DeleteByID() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssignmentsClient().DeleteByID(ctx, "providers/Microsoft.Management/managementGroups/MyManagementGroup/providers/Microsoft.Authorization/policyAssignments/LowCostStorage", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Assignment = armpolicy.Assignment{
	// 	Name: to.Ptr("LowCostStorage"),
	// 	Type: to.Ptr("Microsoft.Authorization/policyAssignments"),
	// 	ID: to.Ptr("/providers/Microsoft.Management/managementGroups/MyManagementGroup/providers/Microsoft.Authorization/policyAssignments/LowCostStorage"),
	// 	Properties: &armpolicy.AssignmentProperties{
	// 		Description: to.Ptr("Allow only storage accounts of SKU Standard_GRS or Standard_LRS to be created"),
	// 		DefinitionVersion: to.Ptr("1.*.*"),
	// 		DisplayName: to.Ptr("Enforce storage account SKU"),
	// 		Metadata: map[string]any{
	// 			"assignedBy": "Cheapskate Boss",
	// 		},
	// 		NotScopes: []*string{
	// 		},
	// 		Parameters: map[string]*armpolicy.ParameterValuesValue{
	// 			"listOfAllowedSKUs": &armpolicy.ParameterValuesValue{
	// 				Value: []any{
	// 					"Standard_GRS",
	// 					"Standard_LRS",
	// 				},
	// 			},
	// 		},
	// 		PolicyDefinitionID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/7433c107-6db4-4ad1-b57a-a76dce0154a1"),
	// 	},
	// }
}
