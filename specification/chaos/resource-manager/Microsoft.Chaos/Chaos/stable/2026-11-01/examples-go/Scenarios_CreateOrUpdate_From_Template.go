package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos/v2"
)

// Generated from example definition: 2026-11-01/Scenarios_CreateOrUpdate_From_Template.json
func ExampleScenariosClient_CreateOrUpdate_createAScenarioFromABuiltInScenarioTemplateVersion() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("6b052e15-03d3-4f17-b2e1-be7f07588291", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScenariosClient().CreateOrUpdate(ctx, "exampleRG", "exampleWorkspace", "computeZoneDownScenario", armchaos.Scenario{
		Properties: &armchaos.ScenarioProperties{
			CreatedFrom: to.Ptr("/providers/Microsoft.Chaos/scenarioTemplates/ComputeZoneDown/versions/1.2"),
			Description: to.Ptr("Simulates a datacenter/availability zone failure by shutting down VMSS instances and standalone VMs in the target zone. Zone failures account for ~25% of all cloud incidents — this scenario validates your application's cross-zone resilience."),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armchaos.ScenariosClientCreateOrUpdateResponse{
	// 	Scenario: armchaos.Scenario{
	// 		ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/workspaces/exampleWorkspace/scenarios/computeZoneDownScenario"),
	// 		Name: to.Ptr("computeZoneDownScenario"),
	// 		Type: to.Ptr("Microsoft.Chaos/workspaces/scenarios"),
	// 		Properties: &armchaos.ScenarioProperties{
	// 			ProvisioningState: to.Ptr(armchaos.ProvisioningStateSucceeded),
	// 			CreatedFrom: to.Ptr("/providers/Microsoft.Chaos/scenarioTemplates/ComputeZoneDown/versions/1.2"),
	// 			Version: to.Ptr("1.0"),
	// 			Description: to.Ptr("Simulates a datacenter/availability zone failure by shutting down VMSS instances and standalone VMs in the target zone. Zone failures account for ~25% of all cloud incidents — this scenario validates your application's cross-zone resilience."),
	// 		},
	// 		SystemData: &armchaos.SystemData{
	// 			CreatedAt: to.Ptr(time.Date(2025, time.January, 15, 10, 30, 0, 0, time.UTC)),
	// 			CreatedBy: to.Ptr("admin@contoso.com"),
	// 			CreatedByType: to.Ptr(armchaos.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2025, time.January, 15, 10, 30, 0, 0, time.UTC)),
	// 			LastModifiedBy: to.Ptr("admin@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armchaos.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
