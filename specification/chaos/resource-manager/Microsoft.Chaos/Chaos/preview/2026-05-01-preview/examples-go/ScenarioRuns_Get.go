package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos/v2"
)

// Generated from example definition: 2026-05-01-preview/ScenarioRuns_Get.json
func ExampleScenarioRunsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("6b052e15-03d3-4f17-b2e1-be7f07588291", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScenarioRunsClient().Get(ctx, "exampleRG", "exampleWorkspace", "12345678-1234-1234-1234-123456789012", "abcd1234-5678-9012-3456-789012345678", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armchaos.ScenarioRunsClientGetResponse{
	// 	ScenarioRun: armchaos.ScenarioRun{
	// 		ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/workspaces/exampleWorkspace/scenarios/exampleScenario/runs/abcd1234-5678-9012-3456-789012345678"),
	// 		Name: to.Ptr("abcd1234-5678-9012-3456-789012345678"),
	// 		Type: to.Ptr("Microsoft.Chaos/workspaces/scenarios/runs"),
	// 		Properties: &armchaos.ScenarioRunProperties{
	// 			ScenarioConfigurationName: to.Ptr("config-5678-9012-3456-789012345678"),
	// 			ScenarioName: to.Ptr("12345678-1234-1234-1234-123456789012"),
	// 			WorkspaceName: to.Ptr("exampleWorkspace"),
	// 			ManagedIdentityPrincipalID: to.Ptr("d04ab567-2c07-43ef-a7f4-4527626b7f56"),
	// 			Status: to.Ptr(armchaos.ScenarioRunStateSucceeded),
	// 			Resources: []*armchaos.ScenarioRunResource{
	// 				{
	// 					ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/exampleVM1"),
	// 				},
	// 				{
	// 					ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/exampleVM2"),
	// 				},
	// 			},
	// 			ScenarioRunJSON: to.Ptr("{\"experiment\":{\"steps\":[{\"name\":\"step1\",\"branches\":[{\"name\":\"branch1\",\"actions\":[{\"name\":\"urn:csci:provider:microsoft:VirtualMachine_Shutdown/1.0\",\"type\":\"continuous\",\"duration\":\"PT10M\",\"selectorId\":\"selector1\"}]}]}]}}"),
	// 			ScenarioRunSummary: []*armchaos.ScenarioRunSummaryAction{
	// 				{
	// 					ActionUrn: to.Ptr("urn:csci:provider:microsoft:VirtualMachine_Shutdown/1.0"),
	// 					State: to.Ptr(armchaos.ScenarioSummaryStateSucceeded),
	// 					StartedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-20T14:30:00.000Z"); return t}()),
	// 					CompletedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-20T14:40:00.000Z"); return t}()),
	// 					Resources: []*armchaos.ScenarioRunResource{
	// 						{
	// 							ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/exampleVM1"),
	// 						},
	// 						{
	// 							ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/exampleVM2"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-20T14:30:00.000Z"); return t}()),
	// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-20T14:40:00.000Z"); return t}()),
	// 			ZoneResolution: &armchaos.ZoneResolutionInfo{
	// 				Mode: to.Ptr(armchaos.ZoneResolutionModePhysical),
	// 				RequestedPhysicalZones: []*string{
	// 					to.Ptr("westus2-az1"),
	// 				},
	// 				SubscriptionZoneMappings: []*armchaos.ZoneResolutionMapping{
	// 					{
	// 						SubscriptionID: to.Ptr("6b052e15-03d3-4f17-b2e1-be7f07588291"),
	// 						ZoneMappings: []*armchaos.PhysicalToLogicalZoneMapping{
	// 							{
	// 								PhysicalZone: to.Ptr("westus2-az1"),
	// 								LogicalZone: to.Ptr("1"),
	// 							},
	// 						},
	// 					},
	// 					{
	// 						SubscriptionID: to.Ptr("a1b2c3d4-5678-9012-3456-abcdef012345"),
	// 						ZoneMappings: []*armchaos.PhysicalToLogicalZoneMapping{
	// 							{
	// 								PhysicalZone: to.Ptr("westus2-az1"),
	// 								LogicalZone: to.Ptr("3"),
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		SystemData: &armchaos.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-20T14:30:00.000Z"); return t}()),
	// 			CreatedBy: to.Ptr("admin@contoso.com"),
	// 			CreatedByType: to.Ptr(armchaos.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-20T14:40:00.000Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("system"),
	// 			LastModifiedByType: to.Ptr(armchaos.CreatedByTypeApplication),
	// 		},
	// 	},
	// }
}
