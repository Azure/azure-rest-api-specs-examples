Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcustomerlockbox%2Farmcustomerlockbox%2Fv0.4.0/sdk/resourcemanager/customerlockbox/armcustomerlockbox/README.md) on how to add the SDK to your project and authenticate.

```go
package armcustomerlockbox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerlockbox/armcustomerlockbox"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/customerlockbox/resource-manager/Microsoft.CustomerLockbox/preview/2018-02-28-preview/examples/Requests_UpdateStatus.json
func ExampleRequestsClient_UpdateStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcustomerlockbox.NewRequestsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.UpdateStatus(ctx,
		"<subscription-id>",
		"<request-id>",
		armcustomerlockbox.Approval{
			Reason: to.Ptr("<reason>"),
			Status: to.Ptr(armcustomerlockbox.StatusApprove),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
