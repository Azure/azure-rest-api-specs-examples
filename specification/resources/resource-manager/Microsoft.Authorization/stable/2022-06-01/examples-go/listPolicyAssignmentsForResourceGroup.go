package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/resources/resource-manager/Microsoft.Authorization/stable/2022-06-01/examples/listPolicyAssignmentsForResourceGroup.json
func ExampleAssignmentsClient_NewListForResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAssignmentsClient().NewListForResourceGroupPager("TestResourceGroup", &armpolicy.AssignmentsClientListForResourceGroupOptions{Filter: to.Ptr("atScope()"),
		Top: nil,
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.AssignmentListResult = armpolicy.AssignmentListResult{
		// 	Value: []*armpolicy.Assignment{
		// 		{
		// 			Name: to.Ptr("TestCostManagement"),
		// 			Type: to.Ptr("Microsoft.Authorization/policyAssignments"),
		// 			ID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/resourceGroups/TestResourceGroup/providers/Microsoft.Authorization/policyAssignments/TestCostManagement"),
		// 			Identity: &armpolicy.Identity{
		// 				Type: to.Ptr(armpolicy.ResourceIdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("e6d23f8d-af97-4fbc-bda6-00604e4e3d0a"),
		// 				TenantID: to.Ptr("4bee2b8a-1bee-47c2-90e9-404241551135"),
		// 			},
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armpolicy.AssignmentProperties{
		// 				Description: to.Ptr("Minimize the risk of accidental cost overruns"),
		// 				DisplayName: to.Ptr("Storage Cost Management"),
		// 				Metadata: map[string]any{
		// 					"category": "Cost Management",
		// 				},
		// 				NotScopes: []*string{
		// 				},
		// 				Parameters: map[string]*armpolicy.ParameterValuesValue{
		// 					"allowedSkus": &armpolicy.ParameterValuesValue{
		// 						Value: "Standard_A1",
		// 					},
		// 				},
		// 				PolicyDefinitionID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyDefinitions/storageSkus"),
		// 				Scope: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/resourceGroups/TestResourceGroup"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("TestTagEnforcement"),
		// 			Type: to.Ptr("Microsoft.Authorization/policyAssignments"),
		// 			ID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/resourceGroups/TestResourceGroup/providers/Microsoft.Authorization/policyAssignments/TestTagEnforcement"),
		// 			Properties: &armpolicy.AssignmentProperties{
		// 				Description: to.Ptr("Ensure a given tag key and value are present on all resources"),
		// 				DisplayName: to.Ptr("Enforces a tag key and value"),
		// 				NotScopes: []*string{
		// 				},
		// 				PolicyDefinitionID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyDefinitions/TagKeyValue"),
		// 				Scope: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/resourceGroups/TestResourceGroup"),
		// 			},
		// 	}},
		// }
	}
}
