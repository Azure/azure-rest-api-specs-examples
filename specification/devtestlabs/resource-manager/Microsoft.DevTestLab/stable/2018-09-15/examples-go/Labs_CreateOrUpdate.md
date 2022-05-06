Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdevtestlabs%2Farmdevtestlabs%2Fv0.4.0/sdk/resourcemanager/devtestlabs/armdevtestlabs/README.md) on how to add the SDK to your project and authenticate.

```go
package armdevtestlabs_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Labs_CreateOrUpdate.json
func ExampleLabsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdevtestlabs.NewLabsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<name>",
		armdevtestlabs.Lab{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"tagName1": to.Ptr("tagValue1"),
			},
			Properties: &armdevtestlabs.LabProperties{
				LabStorageType: to.Ptr(armdevtestlabs.StorageType("{Standard|Premium}")),
			},
		},
		&armdevtestlabs.LabsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
