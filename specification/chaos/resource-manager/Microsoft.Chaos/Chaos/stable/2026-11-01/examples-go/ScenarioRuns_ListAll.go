package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos/v2"
)

// Generated from example definition: 2026-11-01/ScenarioRuns_ListAll.json
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
		// 					ExcludedResources: []*armchaos.ScenarioRunResource{
		// 						{
		// 							ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/exampleVM3"),
		// 						},
		// 					},
		// 					StartTime: to.Ptr(time.Date(2025, time.January, 20, 14, 30, 0, 0, time.UTC)),
		// 					EndTime: to.Ptr(time.Date(2025, time.January, 20, 14, 40, 0, 0, time.UTC)),
		// 				},
		// 				SystemData: &armchaos.SystemData{
		// 					CreatedAt: to.Ptr(time.Date(2025, time.January, 20, 14, 30, 0, 0, time.UTC)),
		// 					CreatedBy: to.Ptr("admin@contoso.com"),
		// 					CreatedByType: to.Ptr(armchaos.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(time.Date(2025, time.January, 20, 14, 40, 0, 0, time.UTC)),
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
		// 					Status: to.Ptr(armchaos.ScenarioRunStatePartiallySucceeded),
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
		// 								ErrorMessage: to.Ptr("The workspace managed identity is missing the required permissions to run this action on the target resource."),
		// 							},
		// 						},
		// 						Resource: []*armchaos.ResourceStateError{
		// 							{
		// 								ResourceID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/exampleVM2"),
		// 								ErrorCode: to.Ptr("ResourceNotReady"),
		// 								ErrorMessage: to.Ptr("The target resource was not ready for fault injection."),
		// 								RemediationURI: to.Ptr("https://aka.ms/azure-chaos-studio"),
		// 							},
		// 						},
		// 					},
		// 					ScenarioRunSummary: []*armchaos.ScenarioRunSummaryAction{
		// 						{
		// 							ActionUrn: to.Ptr("urn:csci:provider:microsoft:VirtualMachine_Restart/1.0"),
		// 							State: to.Ptr(armchaos.ScenarioSummaryStateFailed),
		// 							StartedAt: to.Ptr(time.Date(2025, time.January, 19, 10, 15, 0, 0, time.UTC)),
		// 							CompletedAt: to.Ptr(time.Date(2025, time.January, 19, 10, 18, 0, 0, time.UTC)),
		// 							Resources: []*armchaos.ScenarioRunSummaryResource{
		// 								{
		// 									ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/exampleVM2"),
		// 								},
		// 							},
		// 							Error: &armchaos.ActionError{
		// 								Code: to.Ptr("ActionExecutionFailed"),
		// 								Message: to.Ptr("The action failed on the target resource."),
		// 								Details: []*armchaos.ActionErrorDetail{
		// 									{
		// 										Code: to.Ptr("TargetExecutionFailed"),
		// 										Message: to.Ptr("The target resource was not ready for fault injection."),
		// 										Target: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/exampleVM2"),
		// 									},
		// 								},
		// 							},
		// 							FailureSuppressed: to.Ptr(true),
		// 						},
		// 					},
		// 					Warnings: []*armchaos.ScenarioRunWarning{
		// 						{
		// 							Code: to.Ptr("AgentCleanupIncomplete"),
		// 							Message: to.Ptr("Cleanup could not be confirmed for one target."),
		// 						},
		// 					},
		// 					StartTime: to.Ptr(time.Date(2025, time.January, 19, 10, 15, 0, 0, time.UTC)),
		// 					EndTime: to.Ptr(time.Date(2025, time.January, 19, 10, 18, 0, 0, time.UTC)),
		// 				},
		// 				SystemData: &armchaos.SystemData{
		// 					CreatedAt: to.Ptr(time.Date(2025, time.January, 19, 10, 15, 0, 0, time.UTC)),
		// 					CreatedBy: to.Ptr("admin@contoso.com"),
		// 					CreatedByType: to.Ptr(armchaos.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(time.Date(2025, time.January, 19, 10, 18, 0, 0, time.UTC)),
		// 					LastModifiedBy: to.Ptr("system"),
		// 					LastModifiedByType: to.Ptr(armchaos.CreatedByTypeApplication),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
