Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Ftestbase%2Farmtestbase%2Fv0.1.0/sdk/resourcemanager/testbase/armtestbase/README.md) on how to add the SDK to your project and authenticate.

```go
package armtestbase_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/testbase/armtestbase"
)

// x-ms-original-file: specification/testbase/resource-manager/Microsoft.TestBase/preview/2020-12-16-preview/examples/TestBaseAccountCreate.json
func ExampleTestBaseAccountsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armtestbase.NewTestBaseAccountsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<test-base-account-name>",
		armtestbase.TestBaseAccountResource{
			TrackedResource: armtestbase.TrackedResource{
				Location: to.StringPtr("<location>"),
			},
			Properties: &armtestbase.TestBaseAccountResourceProperties{
				SKU: &armtestbase.TestBaseAccountSKU{
					Name: to.StringPtr("<name>"),
					Tier: armtestbase.TierStandard.ToPtr(),
				},
			},
		},
		&armtestbase.TestBaseAccountsBeginCreateOptions{Restore: nil})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("TestBaseAccountResource.ID: %s\n", *res.ID)
}
```
