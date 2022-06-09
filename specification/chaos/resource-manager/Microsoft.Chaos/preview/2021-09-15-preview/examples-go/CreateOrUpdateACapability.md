```go
package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/chaos/resource-manager/Microsoft.Chaos/preview/2021-09-15-preview/examples/CreateOrUpdateACapability.json
func ExampleCapabilitiesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armchaos.NewCapabilitiesClient("6b052e15-03d3-4f17-b2e1-be7f07588291", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"exampleRG",
		"Microsoft.Compute",
		"virtualMachines",
		"exampleVM",
		"Microsoft-VirtualMachine",
		"Shutdown-1.0",
		armchaos.Capability{
			Properties: &armchaos.CapabilityProperties{},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fchaos%2Farmchaos%2Fv0.5.0/sdk/resourcemanager/chaos/armchaos/README.md) on how to add the SDK to your project and authenticate.
