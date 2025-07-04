package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos/v2"
)

// Generated from example definition: 2025-01-01/TargetTypes_Get.json
func ExampleTargetTypesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("6b052e15-03d3-4f17-b2e1-be7f07588291", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTargetTypesClient().Get(ctx, "westus2", "Microsoft-Agent", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armchaos.TargetTypesClientGetResponse{
	// 	TargetType: &armchaos.TargetType{
	// 		Name: to.Ptr("Microsoft-Agent"),
	// 		Type: to.Ptr("Microsoft.Chaos/locations/targetTypes"),
	// 		ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/providers/Microsoft.Chaos/locations/westus2/targetTypes/Microsoft-Agent"),
	// 		Properties: &armchaos.TargetTypeProperties{
	// 			Description: to.Ptr("A target represents Chaos Agent."),
	// 			DisplayName: to.Ptr("Chaos Agent"),
	// 			PropertiesSchema: to.Ptr("https://schema.centralus.chaos-prod.azure.com/targets/Microsoft-VirtualMachine.json"),
	// 			ResourceTypes: []*string{
	// 				to.Ptr("Microsoft.Compute/virtualMachines"),
	// 				to.Ptr("Microsoft.Compute/virtualMachineScaleSets"),
	// 			},
	// 		},
	// 	},
	// }
}
