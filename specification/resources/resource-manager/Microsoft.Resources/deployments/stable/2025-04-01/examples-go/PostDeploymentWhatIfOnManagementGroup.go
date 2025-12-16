package armdeployments_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armdeployments"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/18609d68cf243ee3ce35d7c005ff3c7dd2cd9477/specification/resources/resource-manager/Microsoft.Resources/deployments/stable/2025-04-01/examples/PostDeploymentWhatIfOnManagementGroup.json
func ExampleDeploymentsClient_BeginWhatIfAtManagementGroupScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeployments.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDeploymentsClient().BeginWhatIfAtManagementGroupScope(ctx, "myManagementGruop", "exampleDeploymentName", armdeployments.ScopedDeploymentWhatIf{
		Location: to.Ptr("eastus"),
		Properties: &armdeployments.DeploymentWhatIfProperties{
			Mode:       to.Ptr(armdeployments.DeploymentModeIncremental),
			Parameters: map[string]*armdeployments.DeploymentParameter{},
			TemplateLink: &armdeployments.TemplateLink{
				URI: to.Ptr("https://example.com/exampleTemplate.json"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WhatIfOperationResult = armdeployments.WhatIfOperationResult{
	// 	Properties: &armdeployments.WhatIfOperationProperties{
	// 		Changes: []*armdeployments.WhatIfChange{
	// 			{
	// 				After: map[string]any{
	// 					"before":map[string]any{
	// 						"name": "myPolicyAssignment",
	// 						"type": "Microsoft.Authorization/policyAssignments",
	// 						"apiVersion": "2019-06-01",
	// 						"id": "/providers/Microsoft.Management/managementGroups/myManagementGroup/providers/Microsoft.Authorization/policyAssignments/myPolicyAssignment",
	// 						"location": "westus2",
	// 						"properties":map[string]any{
	// 							"enforcementMode": "DoNotEnforce",
	// 							"policyDefinitionId": "/providers/Microsoft.Management/managementGroups/myManagementGroup/providers/Microsoft.Authorization/policyAssignments/myPolicyDefinition",
	// 							"scope": "/subscriptions/00000000-0000-0000-0000-000000000001",
	// 						},
	// 					},
	// 					"changeType": "Modify",
	// 					"delta":[]any{
	// 						map[string]any{
	// 							"path": "properties.enforcementMode",
	// 							"after": "DoNotEnforce",
	// 							"before": "Default",
	// 							"propertyChangeType": "Modify",
	// 						},
	// 					},
	// 					"resourceId": "/providers/Microsoft.Management/managementGroups/myManagementGroup/providers/Microsoft.Authorization/policyAssignments/myPolicyAssignment",
	// 				},
	// 				Before: map[string]any{
	// 					"name": "myPolicyAssignment",
	// 					"type": "Microsoft.Authorization/policyAssignments",
	// 					"apiVersion": "2019-06-01",
	// 					"id": "/providers/Microsoft.Management/managementGroups/myManagementGroup/providers/Microsoft.Authorization/policyAssignments/myPolicyAssignment",
	// 					"location": "westus2",
	// 					"properties":map[string]any{
	// 						"enforcementMode": "Default",
	// 						"policyDefinitionId": "/providers/Microsoft.Management/managementGroups/myManagementGroup/providers/Microsoft.Authorization/policyAssignments/myPolicyDefinition",
	// 						"scope": "/subscriptions/00000000-0000-0000-0000-000000000001",
	// 					},
	// 				},
	// 				ChangeType: to.Ptr(armdeployments.ChangeTypeModify),
	// 				ResourceID: to.Ptr("/providers/Microsoft.Management/managementGroups/myManagementGroup/providers/Microsoft.Authorization/policyAssignments/myPolicyAssignment"),
	// 			},
	// 			{
	// 				After: map[string]any{
	// 					"name": "myPolicyAssignment2",
	// 					"type": "Microsoft.Authorization/policyAssignments",
	// 					"apiVersion": "2019-06-01",
	// 					"id": "/providers/Microsoft.Management/managementGroups/myManagementGroup/providers/Microsoft.Authorization/policyAssignments/myPolicyAssignment2",
	// 					"location": "westus2",
	// 					"properties":map[string]any{
	// 						"enforcementMode": "Default",
	// 						"policyDefinitionId": "/providers/Microsoft.Management/managementGroups/myManagementGroup/providers/Microsoft.Authorization/policyAssignments/myPolicyDefinition",
	// 						"scope": "/subscriptions/00000000-0000-0000-0000-000000000002",
	// 					},
	// 				},
	// 				ChangeType: to.Ptr(armdeployments.ChangeTypeCreate),
	// 				ResourceID: to.Ptr("/providers/Microsoft.Management/managementGroups/myManagementGroup/providers/Microsoft.Authorization/policyAssignments/myPolicyAssignment2"),
	// 		}},
	// 	},
	// 	Status: to.Ptr("Succeeded"),
	// }
}
