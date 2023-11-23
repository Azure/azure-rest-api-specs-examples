package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/chaos/resource-manager/Microsoft.Chaos/stable/2023-11-01/examples/GetCapability.json
func ExampleCapabilitiesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCapabilitiesClient().Get(ctx, "exampleRG", "Microsoft.Compute", "virtualMachines", "exampleVM", "Microsoft-VirtualMachine", "Shutdown-1.0", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Capability = armchaos.Capability{
	// 	Name: to.Ptr("Shutdown-1.0"),
	// 	Type: to.Ptr("Microsoft.Chaos/targets/capabilities"),
	// 	ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/exampleVM/providers/Microsoft.Chaos/targets/Microsoft-VirtualMachine/capabilities/Shutdown-1.0"),
	// 	Properties: &armchaos.CapabilityProperties{
	// 		Description: to.Ptr("Shutdown an Azure Virtual Machine for a defined period of time."),
	// 		ParametersSchema: to.Ptr("https://schema.centralus.chaos-prod.azure.com/targets/Microsoft-VirtualMachine/capabilities/Shutdown-1.0.json"),
	// 		Publisher: to.Ptr("Microsoft"),
	// 		TargetType: to.Ptr("VirtualMachine"),
	// 		Urn: to.Ptr("urn:csci:microsoft:virtualMachine:shutdown/1.0"),
	// 	},
	// 	SystemData: &armchaos.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-14T05:08:38.466Z"); return t}()),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-14T05:08:38.466Z"); return t}()),
	// 	},
	// }
}
