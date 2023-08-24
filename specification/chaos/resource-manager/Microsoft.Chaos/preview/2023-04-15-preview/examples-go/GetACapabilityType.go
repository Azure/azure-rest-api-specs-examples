package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/chaos/resource-manager/Microsoft.Chaos/preview/2023-04-15-preview/examples/GetACapabilityType.json
func ExampleCapabilityTypesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCapabilityTypesClient().Get(ctx, "westus2", "Microsoft-VirtualMachine", "Shutdown-1.0", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CapabilityType = armchaos.CapabilityType{
	// 	Name: to.Ptr("Shutdown-1.0"),
	// 	Type: to.Ptr("Microsoft.Chaos/locations/targetTypes/capabilityTypes"),
	// 	ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/Microsoft.Chaos/locations/westus2/targetTypes/Microsoft-VirtualMachine/capabilityTypes/Shutdown-1.0"),
	// 	Properties: &armchaos.CapabilityTypeProperties{
	// 		Description: to.Ptr("Shutdown an Azure Virtual Machine for a defined period of time."),
	// 		DisplayName: to.Ptr("Shutdown VM"),
	// 		Kind: to.Ptr("fault"),
	// 		ParametersSchema: to.Ptr("https://schema.centralus.chaos-prod.azure.com/targets/Microsoft-VirtualMachine/capabilities/Shutdown-1.0.json"),
	// 		Publisher: to.Ptr("Microsoft"),
	// 		RuntimeProperties: &armchaos.CapabilityTypePropertiesRuntimeProperties{
	// 			Kind: to.Ptr("continuous"),
	// 		},
	// 		TargetType: to.Ptr("VirtualMachine"),
	// 		Urn: to.Ptr("urn:csci:microsoft:virtualMachine:shutdown/1.0"),
	// 	},
	// }
}
