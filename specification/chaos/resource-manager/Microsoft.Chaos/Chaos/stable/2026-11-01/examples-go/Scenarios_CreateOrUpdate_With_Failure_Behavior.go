package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos/v2"
)

// Generated from example definition: 2026-11-01/Scenarios_CreateOrUpdate_With_Failure_Behavior.json
func ExampleScenariosClient_CreateOrUpdate_createOrUpdateAScenarioThatContinuesPastSuppressedActionFailures() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("6b052e15-03d3-4f17-b2e1-be7f07588291", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScenariosClient().CreateOrUpdate(ctx, "exampleRG", "exampleWorkspace", "bestEffortProbeScenario", armchaos.Scenario{
		Properties: &armchaos.ScenarioProperties{
			Description:     to.Ptr("Probes dependencies best-effort, then runs a critical shutdown whose failure must stop the run."),
			FailureBehavior: to.Ptr(armchaos.FailureBehaviorContinue),
			Actions: []*armchaos.ScenarioAction{
				{
					Name:        to.Ptr("dependencyProbe"),
					ActionID:    to.Ptr("urn:csci:microsoft:agent:networkDisconnect/1.0.0"),
					Description: to.Ptr("Best-effort dependency probe. Inherits the scenario-level Continue behavior, so a failure here does not fail the run."),
					Duration:    to.Ptr("PT5M"),
				},
				{
					Name:            to.Ptr("criticalShutdown"),
					ActionID:        to.Ptr("urn:csci:microsoft:compute:shutdown/1.0.0"),
					Description:     to.Ptr("Critical shutdown. Overrides the scenario default so that a failure stops the run."),
					Duration:        to.Ptr("PT10M"),
					FailureBehavior: to.Ptr(armchaos.FailureBehaviorStop),
					RunAfter: &armchaos.RunAfter{
						Behavior: to.Ptr(armchaos.RunAfterBehaviorAll),
						Items: []*armchaos.ActionDependency{
							{
								Type:              to.Ptr(armchaos.ActionDependencyTypeAction),
								Name:              to.Ptr("dependencyProbe"),
								OnActionLifecycle: to.Ptr(armchaos.ActionLifecycleAnyTerminal),
							},
						},
					},
				},
			},
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
	// 		ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/workspaces/exampleWorkspace/scenarios/bestEffortProbeScenario"),
	// 		Name: to.Ptr("bestEffortProbeScenario"),
	// 		Type: to.Ptr("Microsoft.Chaos/workspaces/scenarios"),
	// 		Properties: &armchaos.ScenarioProperties{
	// 			ProvisioningState: to.Ptr(armchaos.ProvisioningStateSucceeded),
	// 			Version: to.Ptr("1.0"),
	// 			Description: to.Ptr("Probes dependencies best-effort, then runs a critical shutdown whose failure must stop the run."),
	// 			FailureBehavior: to.Ptr(armchaos.FailureBehaviorContinue),
	// 			Actions: []*armchaos.ScenarioAction{
	// 				{
	// 					Name: to.Ptr("dependencyProbe"),
	// 					ActionID: to.Ptr("urn:csci:microsoft:agent:networkDisconnect/1.0.0"),
	// 					Description: to.Ptr("Best-effort dependency probe. Inherits the scenario-level Continue behavior, so a failure here does not fail the run."),
	// 					Duration: to.Ptr("PT5M"),
	// 				},
	// 				{
	// 					Name: to.Ptr("criticalShutdown"),
	// 					ActionID: to.Ptr("urn:csci:microsoft:compute:shutdown/1.0.0"),
	// 					Description: to.Ptr("Critical shutdown. Overrides the scenario default so that a failure stops the run."),
	// 					Duration: to.Ptr("PT10M"),
	// 					FailureBehavior: to.Ptr(armchaos.FailureBehaviorStop),
	// 					RunAfter: &armchaos.RunAfter{
	// 						Behavior: to.Ptr(armchaos.RunAfterBehaviorAll),
	// 						Items: []*armchaos.ActionDependency{
	// 							{
	// 								Type: to.Ptr(armchaos.ActionDependencyTypeAction),
	// 								Name: to.Ptr("dependencyProbe"),
	// 								OnActionLifecycle: to.Ptr(armchaos.ActionLifecycleAnyTerminal),
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
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
