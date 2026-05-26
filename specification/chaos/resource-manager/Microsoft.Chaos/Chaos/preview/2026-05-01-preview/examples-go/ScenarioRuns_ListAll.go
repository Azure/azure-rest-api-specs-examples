package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos/v2"
)

// Generated from example definition: 2026-05-01-preview/ScenarioRuns_ListAll.json
func ExampleScenarioRunsClient_NewListAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("6b052e15-03d3-4f17-b2e1-be7f07588291", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewScenarioRunsClient().NewListAllPager("exampleRG", "exampleWorkspace", "12345678-1234-1234-1234-123456789012", nil)
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
		// page = armchaos.ScenarioRunsClientListAllResponse{
		// 	ScenarioRunListResult: armchaos.ScenarioRunListResult{
		// 		Value: []*armchaos.ScenarioRun{
		// 			{
		// 				ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/workspaces/exampleWorkspace/scenarios/12345678-1234-1234-1234-123456789012/runs/abcd1234-5678-9012-3456-789012345678"),
		// 				Name: to.Ptr("abcd1234-5678-9012-3456-789012345678"),
		// 				Type: to.Ptr("Microsoft.Chaos/workspaces/scenarios/runs"),
		// 				Properties: &armchaos.ScenarioRunProperties{
		// 					ScenarioConfigurationName: to.Ptr("config-5678-9012-3456-789012345678"),
		// 					ScenarioName: to.Ptr("12345678-1234-1234-1234-123456789012"),
		// 					WorkspaceName: to.Ptr("exampleWorkspace"),
		// 					ManagedIdentityPrincipalID: to.Ptr("d04ab567-2c07-43ef-a7f4-4527626b7f56"),
		// 					Status: to.Ptr(armchaos.ScenarioRunStateSucceeded),
		// 					Resources: []*armchaos.ScenarioRunResource{
		// 						{
		// 							ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/exampleVM1"),
		// 						},
		// 					},
		// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-20T14:30:00.000Z"); return t}()),
		// 					EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-20T14:40:00.000Z"); return t}()),
		// 				},
		// 				SystemData: &armchaos.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-20T14:30:00.000Z"); return t}()),
		// 					CreatedBy: to.Ptr("admin@contoso.com"),
		// 					CreatedByType: to.Ptr(armchaos.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-20T14:40:00.000Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("system"),
		// 					LastModifiedByType: to.Ptr(armchaos.CreatedByTypeApplication),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/workspaces/exampleWorkspace/scenarios/exampleScenario/runs/efgh5678-9012-3456-7890-123456789012"),
		// 				Name: to.Ptr("efgh5678-9012-3456-7890-123456789012"),
		// 				Type: to.Ptr("Microsoft.Chaos/workspaces/scenarios/runs"),
		// 				Properties: &armchaos.ScenarioRunProperties{
		// 					ScenarioConfigurationName: to.Ptr("config-5678-9012-3456-789012345678"),
		// 					ScenarioName: to.Ptr("12345678-1234-1234-1234-123456789012"),
		// 					WorkspaceName: to.Ptr("exampleWorkspace"),
		// 					ManagedIdentityPrincipalID: to.Ptr("d04ab567-2c07-43ef-a7f4-4527626b7f56"),
		// 					Status: to.Ptr(armchaos.ScenarioRunStateFailed),
		// 					Resources: []*armchaos.ScenarioRunResource{
		// 						{
		// 							ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/exampleVM2"),
		// 						},
		// 					},
		// 					ExecutionErrors: &armchaos.ScenarioErrors{
		// 						Permission: []*armchaos.PermissionError{
		// 							{
		// 								ResourceID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/exampleVM2"),
		// 								MissingPermissions: []*string{
		// 									to.Ptr("Microsoft.Compute/virtualMachines/restart/action"),
		// 								},
		// 								RequiredPermissions: []*string{
		// 									to.Ptr("Microsoft.Compute/virtualMachines/restart/action"),
		// 									to.Ptr("Microsoft.Compute/virtualMachines/read"),
		// 								},
		// 								RecommendedRoles: []*string{
		// 									to.Ptr("Virtual Machine Contributor"),
		// 								},
		// 								Identity: &armchaos.EntraIdentity{
		// 									ObjectID: to.Ptr("d04ab567-2c07-43ef-a7f4-4527626b7f56"),
		// 									TenantID: to.Ptr("8c3e2fb2-fe7a-4bf1-b779-d73990782fe6"),
		// 								},
		// 							},
		// 						},
		// 						Resource: []*armchaos.ResourceStateError{
		// 						},
		// 					},
		// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-19T10:15:00.000Z"); return t}()),
		// 					EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-19T10:18:00.000Z"); return t}()),
		// 				},
		// 				SystemData: &armchaos.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-19T10:15:00.000Z"); return t}()),
		// 					CreatedBy: to.Ptr("admin@contoso.com"),
		// 					CreatedByType: to.Ptr(armchaos.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-19T10:18:00.000Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("system"),
		// 					LastModifiedByType: to.Ptr(armchaos.CreatedByTypeApplication),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
