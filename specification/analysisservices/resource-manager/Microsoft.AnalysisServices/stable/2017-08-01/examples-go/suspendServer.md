Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fanalysisservices%2Farmanalysisservices%2Fv0.4.0/sdk/resourcemanager/analysisservices/armanalysisservices/README.md) on how to add the SDK to your project and authenticate.

```go
package armanalysisservices_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/analysisservices/armanalysisservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/analysisservices/resource-manager/Microsoft.AnalysisServices/stable/2017-08-01/examples/suspendServer.json
func ExampleServersClient_BeginSuspend() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armanalysisservices.NewServersClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginSuspend(ctx,
		"<resource-group-name>",
		"<server-name>",
		&armanalysisservices.ServersClientBeginSuspendOptions{ResumeToken: ""})
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
