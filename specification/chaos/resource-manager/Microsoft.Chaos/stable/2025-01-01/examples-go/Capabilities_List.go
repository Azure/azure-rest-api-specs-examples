package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos/v2"
)

// Generated from example definition: 2025-01-01/Capabilities_List.json
func ExampleCapabilitiesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("6b052e15-03d3-4f17-b2e1-be7f07588291", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCapabilitiesClient().NewListPager("exampleRG", "Microsoft.Compute", "virtualMachines", "exampleVM", "Microsoft-VirtualMachine", nil)
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
		// page = armchaos.CapabilitiesClientListResponse{
		// 	CapabilityListResult: armchaos.CapabilityListResult{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/exampleVM/providers/Microsoft.Chaos/targets/Microsoft-VirtualMachine/capabilities?continuationToken=&api-version=2024-11-01-preview"),
		// 		Value: []*armchaos.Capability{
		// 			{
		// 				Name: to.Ptr("Shutdown-1.0"),
		// 				Type: to.Ptr("Microsoft.Chaos/targets/capabilities"),
		// 				ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/exampleVM/providers/Microsoft.Chaos/targets/Microsoft-VirtualMachine/capabilities/Shutdown-1.0"),
		// 				Properties: &armchaos.CapabilityProperties{
		// 					Description: to.Ptr("Shutdown an Azure Virtual Machine for a defined period of time."),
		// 					ParametersSchema: to.Ptr("https://schema.centralus.chaos-prod.azure.com/targets/Microsoft-VirtualMachine/capabilities/Shutdown-1.0.json"),
		// 					Publisher: to.Ptr("Microsoft"),
		// 					TargetType: to.Ptr("VirtualMachine"),
		// 					Urn: to.Ptr("urn:csci:microsoft:virtualMachine:shutdown/1.0"),
		// 				},
		// 				SystemData: &armchaos.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-14T05:08:38.4662189Z"); return t}()),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-14T05:08:38.4662189Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
