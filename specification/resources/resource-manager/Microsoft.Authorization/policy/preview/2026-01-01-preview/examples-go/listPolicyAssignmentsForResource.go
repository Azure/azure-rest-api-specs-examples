package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy/v2"
)

// Generated from example definition: 2026-01-01-preview/listPolicyAssignmentsForResource.json
func ExampleAssignmentsClient_NewListForResourcePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("ae640e6b-ba3e-4256-9d62-2993eecfa6f2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAssignmentsClient().NewListForResourcePager("TestResourceGroup", "Microsoft.Compute", "virtualMachines/MyTestVm", "domainNames", "MyTestComputer.cloudapp.net", nil)
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
		// page = armpolicy.AssignmentsClientListForResourceResponse{
		// 	AssignmentListResult: armpolicy.AssignmentListResult{
		// 		Value: []*armpolicy.Assignment{
		// 			{
		// 				ID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/resourceGroups/TestResourceGroup/providers/Microsoft.Authorization/policyAssignments/TestCostManagement"),
		// 				Type: to.Ptr("Microsoft.Authorization/policyAssignments"),
		// 				Name: to.Ptr("TestCostManagement"),
		// 				Location: to.Ptr("eastus"),
		// 				Identity: &armpolicy.Identity{
		// 					Type: to.Ptr(armpolicy.ResourceIdentityTypeSystemAssigned),
		// 					PrincipalID: to.Ptr("e6d23f8d-af97-4fbc-bda6-00604e4e3d0a"),
		// 					TenantID: to.Ptr("4bee2b8a-1bee-47c2-90e9-404241551135"),
		// 				},
		// 				Properties: &armpolicy.AssignmentProperties{
		// 					DisplayName: to.Ptr("VM Cost Management"),
		// 					Description: to.Ptr("Minimize the risk of accidental cost overruns"),
		// 					Metadata: map[string]any{
		// 						"category": "Cost Management",
		// 					},
		// 					PolicyDefinitionID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyDefinitions/vmSkus"),
		// 					DefinitionVersion: to.Ptr("1.*.*"),
		// 					Parameters: map[string]*armpolicy.ParameterValuesValue{
		// 						"allowedSkus": &armpolicy.ParameterValuesValue{
		// 							Value: "Standard_A1",
		// 						},
		// 					},
		// 					Scope: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/resourceGroups/TestResourceGroup"),
		// 					NotScopes: []*string{
		// 					},
		// 					InstanceID: to.Ptr("e9a0b1c2-d3e4-5f6a-7b8c-9d0e1f2a3b4c"),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/resourceGroups/TestResourceGroup/providers/Microsoft.Authorization/policyAssignments/TestTagEnforcement"),
		// 				Type: to.Ptr("Microsoft.Authorization/policyAssignments"),
		// 				Name: to.Ptr("TestTagEnforcement"),
		// 				Properties: &armpolicy.AssignmentProperties{
		// 					DisplayName: to.Ptr("Enforces a tag key and value"),
		// 					Description: to.Ptr("Ensure a given tag key and value are present on all resources"),
		// 					PolicyDefinitionID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyDefinitions/TagKeyValue"),
		// 					DefinitionVersion: to.Ptr("1.*.*"),
		// 					Scope: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/resourceGroups/TestResourceGroup"),
		// 					NotScopes: []*string{
		// 					},
		// 					InstanceID: to.Ptr("f0b1c2d3-e4f5-6a7b-8c9d-0e1f2a3b4c5d"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
