```go
package armscvmm_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/scvmm/armscvmm"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/CreateCheckpointVirtualMachine.json
func ExampleVirtualMachinesClient_BeginCreateCheckpoint() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armscvmm.NewVirtualMachinesClient("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateCheckpoint(ctx,
		"testrg",
		"DemoVM",
		&armscvmm.VirtualMachinesClientBeginCreateCheckpointOptions{Body: &armscvmm.VirtualMachineCreateCheckpoint{
			Name:        to.Ptr("Demo Checkpoint name"),
			Description: to.Ptr("Demo Checkpoint description"),
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fscvmm%2Farmscvmm%2Fv0.2.0/sdk/resourcemanager/scvmm/armscvmm/README.md) on how to add the SDK to your project and authenticate.
