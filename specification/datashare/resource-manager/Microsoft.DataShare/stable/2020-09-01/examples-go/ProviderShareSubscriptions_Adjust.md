```go
package armdatashare_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/ProviderShareSubscriptions_Adjust.json
func ExampleProviderShareSubscriptionsClient_Adjust() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatashare.NewProviderShareSubscriptionsClient("12345678-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Adjust(ctx,
		"SampleResourceGroup",
		"Account1",
		"Share1",
		"4256e2cf-0f82-4865-961b-12f83333f487",
		armdatashare.ProviderShareSubscription{
			Properties: &armdatashare.ProviderShareSubscriptionProperties{
				ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-26T22:33:24.5785265Z"); return t }()),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdatashare%2Farmdatashare%2Fv1.0.0/sdk/resourcemanager/datashare/armdatashare/README.md) on how to add the SDK to your project and authenticate.
