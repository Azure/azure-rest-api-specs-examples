Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsaas%2Farmsaas%2Fv0.2.0/sdk/resourcemanager/saas/armsaas/README.md) on how to add the SDK to your project and authenticate.

```go
package armsaas_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/saas/armsaas"
)

// x-ms-original-file: specification/saas/resource-manager/Microsoft.SaaS/preview/2018-03-01-beta/examples/saasSubscriptionLevel/SaasUnsubscribe.json
func ExampleSubscriptionLevelClient_BeginUpdateToUnsubscribed() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsaas.NewSubscriptionLevelClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdateToUnsubscribed(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armsaas.DeleteOptions{
			Feedback:        to.StringPtr("<feedback>"),
			ReasonCode:      to.Float32Ptr(0),
			UnsubscribeOnly: to.BoolPtr(true),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
```
