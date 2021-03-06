package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace"
)

// x-ms-original-file: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-06-01/examples/AcknowledgeNotification.json
func ExamplePrivateStoreClient_AcknowledgeOfferNotification() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplace.NewPrivateStoreClient(cred, nil)
	_, err = client.AcknowledgeOfferNotification(ctx,
		"<private-store-id>",
		"<offer-id>",
		&armmarketplace.PrivateStoreClientAcknowledgeOfferNotificationOptions{Payload: &armmarketplace.AcknowledgeOfferNotificationProperties{
			Properties: &armmarketplace.AcknowledgeOfferNotificationDetails{
				Acknowledge: to.BoolPtr(false),
				Dismiss:     to.BoolPtr(false),
				RemoveOffer: to.BoolPtr(false),
				RemovePlans: []*string{
					to.StringPtr("testPlanA")},
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
}
