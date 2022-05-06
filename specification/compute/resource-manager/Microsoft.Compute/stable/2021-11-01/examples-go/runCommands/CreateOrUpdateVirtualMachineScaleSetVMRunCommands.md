Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcompute%2Farmcompute%2Fv0.7.0/sdk/resourcemanager/compute/armcompute/README.md) on how to add the SDK to your project and authenticate.

```go
package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/runCommands/CreateOrUpdateVirtualMachineScaleSetVMRunCommands.json
func ExampleVirtualMachineScaleSetVMRunCommandsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetVMRunCommandsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		"<instance-id>",
		"<run-command-name>",
		armcompute.VirtualMachineRunCommand{
			Location: to.Ptr("<location>"),
			Properties: &armcompute.VirtualMachineRunCommandProperties{
				AsyncExecution: to.Ptr(false),
				Parameters: []*armcompute.RunCommandInputParameter{
					{
						Name:  to.Ptr("<name>"),
						Value: to.Ptr("<value>"),
					},
					{
						Name:  to.Ptr("<name>"),
						Value: to.Ptr("<value>"),
					}},
				RunAsPassword: to.Ptr("<run-as-password>"),
				RunAsUser:     to.Ptr("<run-as-user>"),
				Source: &armcompute.VirtualMachineRunCommandScriptSource{
					Script: to.Ptr("<script>"),
				},
				TimeoutInSeconds: to.Ptr[int32](3600),
			},
		},
		&armcompute.VirtualMachineScaleSetVMRunCommandsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
