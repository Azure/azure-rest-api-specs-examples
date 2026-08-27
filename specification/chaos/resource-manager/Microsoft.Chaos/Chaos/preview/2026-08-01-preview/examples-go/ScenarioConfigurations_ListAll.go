package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos/v3"
)

// Generated from example definition: 2026-08-01-preview/ScenarioConfigurations_ListAll.json
func ExampleScenarioConfigurationsClient_NewListAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("6b052e15-03d3-4f17-b2e1-be7f07588291", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewScenarioConfigurationsClient().NewListAllPager("exampleRG", "exampleWorkspace", "12345678-1234-1234-1234-123456789012", nil)
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
		// page = armchaos.ScenarioConfigurationsClientListAllResponse{
		// 	ScenarioConfigurationListResult: armchaos.ScenarioConfigurationListResult{
		// 		Value: []*armchaos.ScenarioConfiguration{
		// 			{
		// 				ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/workspaces/exampleWorkspace/scenarios/12345678-1234-1234-1234-123456789012/configurations/config-5678-9012-3456-789012345678"),
		// 				Name: to.Ptr("config-5678-9012-3456-789012345678"),
		// 				Type: to.Ptr("Microsoft.Chaos/workspaces/scenarios/configurations"),
		// 				Properties: &armchaos.ScenarioConfigurationProperties{
		// 					ScenarioID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/workspaces/exampleWorkspace/scenarios/12345678-1234-1234-1234-123456789012"),
		// 					Parameters: []*armchaos.KeyValuePair{
		// 						{
		// 							Key: to.Ptr("duration"),
		// 							Value: to.Ptr("PT10M"),
		// 						},
		// 						{
		// 							Key: to.Ptr("targetResourceIds"),
		// 							Value: to.Ptr("[\"/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/vm1\",\"/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/vm2\"]"),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armchaos.ProvisioningStateSucceeded),
		// 					ResourceTargeting: &armchaos.ResourceTargeting{
		// 						Exclude: &armchaos.ResourceTargetingCriteria{
		// 							Resources: []*string{
		// 								to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/protectedVM"),
		// 							},
		// 							Tags: []*armchaos.KeyValuePair{
		// 								{
		// 									Key: to.Ptr("environment"),
		// 									Value: to.Ptr("production"),
		// 								},
		// 							},
		// 							Types: []*string{
		// 								to.Ptr("Microsoft.Compute/virtualMachineScaleSets"),
		// 							},
		// 						},
		// 					},
		// 				},
		// 				SystemData: &armchaos.SystemData{
		// 					CreatedAt: to.Ptr(time.Date(2025, time.January, 15, 10, 30, 0, 0, time.UTC)),
		// 					CreatedBy: to.Ptr("admin@contoso.com"),
		// 					CreatedByType: to.Ptr(armchaos.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(time.Date(2025, time.January, 15, 12, 0, 0, 0, time.UTC)),
		// 					LastModifiedBy: to.Ptr("admin@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armchaos.CreatedByTypeUser),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/workspaces/exampleWorkspace/scenarios/12345678-1234-1234-1234-123456789012/configurations/config-abcd-1234-5678-901234567890"),
		// 				Name: to.Ptr("config-abcd-1234-5678-901234567890"),
		// 				Type: to.Ptr("Microsoft.Chaos/workspaces/scenarios/configurations"),
		// 				Properties: &armchaos.ScenarioConfigurationProperties{
		// 					ScenarioID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/workspaces/exampleWorkspace/scenarios/12345678-1234-1234-1234-123456789012"),
		// 					Parameters: []*armchaos.KeyValuePair{
		// 						{
		// 							Key: to.Ptr("duration"),
		// 							Value: to.Ptr("PT5M"),
		// 						},
		// 						{
		// 							Key: to.Ptr("targetResourceIds"),
		// 							Value: to.Ptr("[\"/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/vm3\"]"),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armchaos.ProvisioningStateCreating),
		// 				},
		// 				SystemData: &armchaos.SystemData{
		// 					CreatedAt: to.Ptr(time.Date(2025, time.January, 20, 9, 0, 0, 0, time.UTC)),
		// 					CreatedBy: to.Ptr("admin@contoso.com"),
		// 					CreatedByType: to.Ptr(armchaos.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(time.Date(2025, time.January, 20, 9, 0, 0, 0, time.UTC)),
		// 					LastModifiedBy: to.Ptr("admin@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armchaos.CreatedByTypeUser),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
