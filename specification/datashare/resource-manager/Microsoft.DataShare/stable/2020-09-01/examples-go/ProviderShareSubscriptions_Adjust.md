Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdatashare%2Farmdatashare%2Fv0.1.0/sdk/resourcemanager/datashare/armdatashare/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/ProviderShareSubscriptions_Adjust.json
func ExampleProviderShareSubscriptionsClient_Adjust() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatashare.NewProviderShareSubscriptionsClient("<subscription-id>", cred, nil)
	res, err := client.Adjust(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<share-name>",
		"<provider-share-subscription-id>",
		armdatashare.ProviderShareSubscription{
			Properties: &armdatashare.ProviderShareSubscriptionProperties{
				ExpirationDate: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-26T22:33:24.5785265Z"); return t }()),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ProviderShareSubscription.ID: %s\n", *res.ID)
}
```
