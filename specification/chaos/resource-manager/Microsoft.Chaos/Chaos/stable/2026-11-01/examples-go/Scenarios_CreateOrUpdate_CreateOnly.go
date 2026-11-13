package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos/v2"
)

// Generated from example definition: 2026-11-01/Scenarios_CreateOrUpdate_CreateOnly.json
func ExampleScenariosClient_CreateOrUpdate_createAScenarioOnlyWhenAbsentOrReturn412WhenItAlreadyExists() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("6b052e15-03d3-4f17-b2e1-be7f07588291", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewScenariosClient().CreateOrUpdate(ctx, "exampleRG", "exampleWorkspace", "exampleScenario", armchaos.Scenario{
		Properties: &armchaos.ScenarioProperties{
			Parameters: []*armchaos.ScenarioParameter{},
			Actions: []*armchaos.ScenarioAction{
				{
					Name:     to.Ptr("restart"),
					ActionID: to.Ptr("microsoft-compute-shutdown/2.0"),
					Duration: to.Ptr("PT5M"),
					Selectors: []*armchaos.ScenarioActionSelector{
						{
							Name:  to.Ptr(armchaos.ScenarioActionSelectorNameLogicalZone),
							Value: to.Ptr("1"),
						},
						{
							Name:  to.Ptr(armchaos.ScenarioActionSelectorNameLocation),
							Value: to.Ptr("eastus"),
						},
					},
				},
			},
		},
	}, &armchaos.ScenariosClientCreateOrUpdateOptions{
		IfNoneMatch: to.Ptr("*")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
