package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos/v2"
)

// Generated from example definition: 2026-11-01/ScenarioTemplates_Get.json
func ExampleScenarioTemplatesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScenarioTemplatesClient().Get(ctx, "ComputeZoneDown", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armchaos.ScenarioTemplatesClientGetResponse{
	// 	ScenarioTemplate: armchaos.ScenarioTemplate{
	// 		ID: to.Ptr("/providers/Microsoft.Chaos/scenarioTemplates/ComputeZoneDown"),
	// 		Name: to.Ptr("ComputeZoneDown"),
	// 		Type: to.Ptr("Microsoft.Chaos/scenarioTemplates"),
	// 		Properties: &armchaos.ScenarioTemplateProperties{
	// 			LatestVersion: to.Ptr("1.2"),
	// 			DisplayName: to.Ptr("Compute Zone Down"),
	// 			Description: to.Ptr("Simulates a datacenter/availability zone failure by shutting down VMSS instances and standalone VMs in the target zone. Zone failures account for ~25% of all cloud incidents — this scenario validates your application's cross-zone resilience."),
	// 			Parameters: []*armchaos.ScenarioParameter{
	// 				{
	// 					Name: to.Ptr("physicalZone"),
	// 					Type: to.Ptr(armchaos.ParameterTypeString),
	// 					Required: to.Ptr(true),
	// 					Description: to.Ptr("Full physical availability-zone identifier to test, for example eastus-az1."),
	// 				},
	// 				{
	// 					Name: to.Ptr("duration"),
	// 					Type: to.Ptr(armchaos.ParameterTypeString),
	// 					Default: to.Ptr("PT15M"),
	// 					Required: to.Ptr(false),
	// 					Description: to.Ptr("The duration of the zone down scenario."),
	// 					Format: &armchaos.ScenarioParameterFormat{
	// 						FormatType: to.Ptr(armchaos.ParameterFormatTypeDuration),
	// 					},
	// 				},
	// 			},
	// 			Actions: []*armchaos.ScenarioAction{
	// 				{
	// 					Name: to.Ptr("vmssZoneShutdown"),
	// 					ActionID: to.Ptr("microsoft-compute-shutdown/2.0"),
	// 					Description: to.Ptr("Shut down VMSS instances in target zone"),
	// 					Duration: to.Ptr("%%{parameters.duration}%%"),
	// 					Parameters: []*armchaos.KeyValuePair{
	// 					},
	// 					Selectors: []*armchaos.ScenarioActionSelector{
	// 						{
	// 							Name: to.Ptr(armchaos.ScenarioActionSelectorNamePhysicalZone),
	// 							Value: to.Ptr("%%{parameters.physicalZone}%%"),
	// 						},
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("vmZoneShutdown"),
	// 					ActionID: to.Ptr("microsoft-compute-shutdown/2.0"),
	// 					Description: to.Ptr("Shut down VMs in target zone"),
	// 					Duration: to.Ptr("%%{parameters.duration}%%"),
	// 					Parameters: []*armchaos.KeyValuePair{
	// 					},
	// 					Selectors: []*armchaos.ScenarioActionSelector{
	// 						{
	// 							Name: to.Ptr(armchaos.ScenarioActionSelectorNamePhysicalZone),
	// 							Value: to.Ptr("%%{parameters.physicalZone}%%"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
