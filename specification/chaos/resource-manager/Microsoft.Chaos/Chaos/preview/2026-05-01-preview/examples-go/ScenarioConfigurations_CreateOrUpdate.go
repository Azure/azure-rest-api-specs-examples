package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos/v2"
)

// Generated from example definition: 2026-05-01-preview/ScenarioConfigurations_CreateOrUpdate.json
func ExampleScenarioConfigurationsClient_BeginCreateOrUpdate_createOrUpdateAScenarioConfiguration() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("6b052e15-03d3-4f17-b2e1-be7f07588291", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewScenarioConfigurationsClient().BeginCreateOrUpdate(ctx, "exampleRG", "exampleWorkspace", "12345678-1234-1234-1234-123456789012", "config-5678-9012-3456-789012345678", armchaos.ScenarioConfiguration{
		Properties: &armchaos.ScenarioConfigurationProperties{
			ScenarioID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/workspaces/exampleWorkspace/scenarios/12345678-1234-1234-1234-123456789012"),
			Exclusions: &armchaos.ConfigurationExclusions{
				Resources: []*string{
					to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/protectedVM"),
				},
				Tags: []*armchaos.KeyValuePair{
					{
						Key:   to.Ptr("environment"),
						Value: to.Ptr("production"),
					},
				},
				Types: []*string{
					to.Ptr("Microsoft.Compute/virtualMachineScaleSets"),
				},
			},
			Parameters: []*armchaos.KeyValuePair{
				{
					Key:   to.Ptr("duration"),
					Value: to.Ptr("PT10M"),
				},
				{
					Key:   to.Ptr("targetResourceIds"),
					Value: to.Ptr("[\"/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/vm1\",\"/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/vm2\"]"),
				},
			},
			Filters: &armchaos.ConfigurationFilters{
				Locations: []*string{
					to.Ptr("eastus"),
				},
				Zones: []*string{
					to.Ptr("1"),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armchaos.ScenarioConfigurationsClientCreateOrUpdateResponse{
	// 	ScenarioConfiguration: armchaos.ScenarioConfiguration{
	// 		ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/workspaces/exampleWorkspace/scenarios/12345678-1234-1234-1234-123456789012/configurations/config-5678-9012-3456-789012345678"),
	// 		Name: to.Ptr("config-5678-9012-3456-789012345678"),
	// 		Type: to.Ptr("Microsoft.Chaos/workspaces/scenarios/configurations"),
	// 		Properties: &armchaos.ScenarioConfigurationProperties{
	// 			ScenarioID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/workspaces/exampleWorkspace/scenarios/12345678-1234-1234-1234-123456789012"),
	// 			Exclusions: &armchaos.ConfigurationExclusions{
	// 				Resources: []*string{
	// 					to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/protectedVM"),
	// 				},
	// 				Tags: []*armchaos.KeyValuePair{
	// 					{
	// 						Key: to.Ptr("environment"),
	// 						Value: to.Ptr("production"),
	// 					},
	// 				},
	// 				Types: []*string{
	// 					to.Ptr("Microsoft.Compute/virtualMachineScaleSets"),
	// 				},
	// 			},
	// 			Parameters: []*armchaos.KeyValuePair{
	// 				{
	// 					Key: to.Ptr("duration"),
	// 					Value: to.Ptr("PT10M"),
	// 				},
	// 				{
	// 					Key: to.Ptr("targetResourceIds"),
	// 					Value: to.Ptr("[\"/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/vm1\",\"/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/vm2\"]"),
	// 				},
	// 			},
	// 			Filters: &armchaos.ConfigurationFilters{
	// 				Locations: []*string{
	// 					to.Ptr("eastus"),
	// 				},
	// 				Zones: []*string{
	// 					to.Ptr("1"),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armchaos.ProvisioningStateSucceeded),
	// 		},
	// 		SystemData: &armchaos.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-15T10:30:00.000Z"); return t}()),
	// 			CreatedBy: to.Ptr("admin@contoso.com"),
	// 			CreatedByType: to.Ptr(armchaos.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-15T12:00:00.000Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("admin@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armchaos.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
