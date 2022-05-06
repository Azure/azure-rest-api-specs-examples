Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhardwaresecuritymodules%2Farmhardwaresecuritymodules%2Fv0.4.0/sdk/resourcemanager/hardwaresecuritymodules/armhardwaresecuritymodules/README.md) on how to add the SDK to your project and authenticate.

```go
package armhardwaresecuritymodules_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hardwaresecuritymodules/armhardwaresecuritymodules"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/stable/2021-11-30/examples/DedicatedHsm_Delete.json
func ExampleDedicatedHsmClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armhardwaresecuritymodules.NewDedicatedHsmClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<name>",
		&armhardwaresecuritymodules.DedicatedHsmClientBeginDeleteOptions{ResumeToken: ""})
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
