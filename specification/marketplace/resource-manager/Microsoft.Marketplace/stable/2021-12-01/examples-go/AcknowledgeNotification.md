Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmarketplace%2Farmmarketplace%2Fv0.4.0/sdk/resourcemanager/marketplace/armmarketplace/README.md) on how to add the SDK to your project and authenticate.

```go
package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/AcknowledgeNotification.json
func ExamplePrivateStoreClient_AcknowledgeOfferNotification() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	_, err = client.AcknowledgeOfferNotification(ctx,
		"<private-store-id>",
		"<offer-id>",
		&armmarketplace.PrivateStoreClientAcknowledgeOfferNotificationOptions{Payload: &armmarketplace.AcknowledgeOfferNotificationProperties{
			Properties: &armmarketplace.AcknowledgeOfferNotificationDetails{
				Acknowledge: to.Ptr(false),
				Dismiss:     to.Ptr(false),
				RemoveOffer: to.Ptr(false),
				RemovePlans: []*string{
					to.Ptr("testPlanA")},
			},
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
}
```
