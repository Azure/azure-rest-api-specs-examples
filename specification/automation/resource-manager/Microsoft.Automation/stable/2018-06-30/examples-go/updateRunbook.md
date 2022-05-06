Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fautomation%2Farmautomation%2Fv0.5.0/sdk/resourcemanager/automation/armautomation/README.md) on how to add the SDK to your project and authenticate.

```go
package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/updateRunbook.json
func ExampleRunbookClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armautomation.NewRunbookClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	_, err = client.Update(ctx,
		"<resource-group-name>",
		"<automation-account-name>",
		"<runbook-name>",
		armautomation.RunbookUpdateParameters{
			Properties: &armautomation.RunbookUpdateProperties{
				Description:      to.Ptr("<description>"),
				LogActivityTrace: to.Ptr[int32](1),
				LogProgress:      to.Ptr(true),
				LogVerbose:       to.Ptr(false),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
}
```
