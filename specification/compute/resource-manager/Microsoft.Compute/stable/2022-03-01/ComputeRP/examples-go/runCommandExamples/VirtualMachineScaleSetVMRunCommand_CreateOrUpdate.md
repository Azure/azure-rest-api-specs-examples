Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcompute%2Farmcompute%2Fv1.0.0/sdk/resourcemanager/compute/armcompute/README.md) on how to add the SDK to your project and authenticate.

```go
package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/runCommandExamples/VirtualMachineScaleSetVMRunCommand_CreateOrUpdate.json
func ExampleVirtualMachineScaleSetVMRunCommandsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetVMRunCommandsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myvmScaleSet",
		"0",
		"myRunCommand",
		armcompute.VirtualMachineRunCommand{
			Location: to.Ptr("West US"),
			Properties: &armcompute.VirtualMachineRunCommandProperties{
				AsyncExecution: to.Ptr(false),
				Parameters: []*armcompute.RunCommandInputParameter{
					{
						Name:  to.Ptr("param1"),
						Value: to.Ptr("value1"),
					},
					{
						Name:  to.Ptr("param2"),
						Value: to.Ptr("value2"),
					}},
				RunAsPassword: to.Ptr("<runAsPassword>"),
				RunAsUser:     to.Ptr("user1"),
				Source: &armcompute.VirtualMachineRunCommandScriptSource{
					Script: to.Ptr("Write-Host Hello World!"),
				},
				TimeoutInSeconds: to.Ptr[int32](3600),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
