package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e4009d2f8d3bf0271757e522c7d1c1997e193d44/specification/chaos/resource-manager/Microsoft.Chaos/stable/2024-01-01/examples/ListTargetTypes.json
func ExampleTargetTypesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTargetTypesClient().NewListPager("westus2", &armchaos.TargetTypesClientListOptions{ContinuationToken: nil})
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
		// page.TargetTypeListResult = armchaos.TargetTypeListResult{
		// 	Value: []*armchaos.TargetType{
		// 		{
		// 			Name: to.Ptr("Microsoft-Agent"),
		// 			Type: to.Ptr("Microsoft.Chaos/locations/targetTypes"),
		// 			ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/providers/Microsoft.Chaos/locations/westus2/targetTypes/Microsoft-Agent"),
		// 			Location: to.Ptr("centraluseuap"),
		// 			Properties: &armchaos.TargetTypeProperties{
		// 				Description: to.Ptr("A target represents Chaos Agent."),
		// 				DisplayName: to.Ptr("Chaos Agent"),
		// 				PropertiesSchema: to.Ptr("https://schema.centralus.chaos-prod.azure.com/targets/Microsoft-VirtualMachine.json"),
		// 				ResourceTypes: []*string{
		// 					to.Ptr("Microsoft.Compute/virtualMachines"),
		// 					to.Ptr("Microsoft.Compute/virtualMachineScaleSets")},
		// 				},
		// 		}},
		// 	}
	}
}
