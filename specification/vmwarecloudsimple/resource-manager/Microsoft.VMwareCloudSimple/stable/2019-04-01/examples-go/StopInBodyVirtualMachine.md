Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fvmwarecloudsimple%2Farmvmwarecloudsimple%2Fv0.4.0/sdk/resourcemanager/vmwarecloudsimple/armvmwarecloudsimple/README.md) on how to add the SDK to your project and authenticate.

```go
package armvmwarecloudsimple_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/vmwarecloudsimple/armvmwarecloudsimple"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/StopInBodyVirtualMachine.json
func ExampleVirtualMachinesClient_BeginStop() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armvmwarecloudsimple.NewVirtualMachinesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginStop(ctx,
		"<resource-group-name>",
		"<referer>",
		"<virtual-machine-name>",
		&armvmwarecloudsimple.VirtualMachinesClientBeginStopOptions{Mode: nil,
			M:           &armvmwarecloudsimple.VirtualMachineStopMode{},
			ResumeToken: "",
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}
```
