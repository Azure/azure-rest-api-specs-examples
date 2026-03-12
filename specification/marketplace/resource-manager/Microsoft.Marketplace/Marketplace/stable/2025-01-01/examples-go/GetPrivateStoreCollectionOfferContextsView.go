package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace/v2"
)

// Generated from example definition: 2025-01-01/GetPrivateStoreCollectionOfferContextsView.json
func ExamplePrivateStoreCollectionOfferClient_ContextsView() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmarketplace.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateStoreCollectionOfferClient().ContextsView(ctx, "a0e28e55-90c4-41d8-8e34-bb7ef7775406", "56a1a02d-8cf8-45df-bf37-d5f7120fcb3d", "mktp3pp.kuku-buku", &armmarketplace.PrivateStoreCollectionOfferClientContextsViewOptions{
		Payload: &armmarketplace.CollectionOffersByAllContextsPayload{
			Properties: &armmarketplace.CollectionOffersByAllContextsProperties{
				SubscriptionIDs: []*string{
					to.Ptr("b340914e-353d-453a-85fb-8f9b65b51f91"),
					to.Ptr("f2baa04d-5bfc-461b-b6d8-61b403c9ec48"),
				},
			},
		}})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmarketplace.PrivateStoreCollectionOfferClientContextsViewResponse{
	// 	Offer: &armmarketplace.Offer{
	// 		Name: to.Ptr("mktp3pp.kuku-buku"),
	// 		Type: to.Ptr("Microsoft.Marketplace/privateStores/collections/offers"),
	// 		ID: to.Ptr("/providers/Microsoft.Marketplace/privateStores/a0e28e55-90c4-41d8-8e34-bb7ef7775406/collections/56a1a02d-8cf8-45df-bf37-d5f7120fcb3d/offers/mktp3pp.kuku-buku"),
	// 		Properties: &armmarketplace.OfferProperties{
	// 			CreatedAt: to.Ptr("05/28/2025 5:50"),
	// 			ETag: to.Ptr("\"9301f4fd-0000-0100-0000-5e248b350000\""),
	// 			IconFileUris: map[string]*string{
	// 				"large": to.Ptr("https://store-images.s-microsoft.com/image/apps.3060"),
	// 				"medium": to.Ptr("https://store-images.s-microsoft.com/image/apps.14425"),
	// 				"small": to.Ptr("https://store-images.s-microsoft.com/image/apps.10777"),
	// 			},
	// 			IsStopSell: to.Ptr(false),
	// 			ModifiedAt: to.Ptr("05/29/2025 5:50"),
	// 			OfferDisplayName: to.Ptr("kuku-buku 3pp"),
	// 			Plans: []*armmarketplace.Plan{
	// 				{
	// 					Accessibility: to.Ptr(armmarketplace.AccessibilityPublic),
	// 					IsStopSell: to.Ptr(true),
	// 					PlanDisplayName: to.Ptr("Kuku 2025.1"),
	// 					PlanID: to.Ptr("0001"),
	// 					SKUID: to.Ptr("0001"),
	// 				},
	// 				{
	// 					Accessibility: to.Ptr(armmarketplace.AccessibilityPrivateSubscriptionOnLevel),
	// 					IsStopSell: to.Ptr(false),
	// 					PlanDisplayName: to.Ptr("Kuku 2025.2"),
	// 					PlanID: to.Ptr("0001-arm64"),
	// 					SKUID: to.Ptr("000A"),
	// 				},
	// 				{
	// 					Accessibility: to.Ptr(armmarketplace.AccessibilityPublic),
	// 					IsStopSell: to.Ptr(false),
	// 					PlanDisplayName: to.Ptr("Kuku 2025.3"),
	// 					PlanID: to.Ptr("0002"),
	// 					SKUID: to.Ptr("0002"),
	// 				},
	// 			},
	// 			PrivateStoreID: to.Ptr("a0e28e55-90c4-41d8-8e34-bb7ef7775406"),
	// 			PublisherDisplayName: to.Ptr("Marketplace Third Party"),
	// 			SpecificPlanIDsLimitation: []*string{
	// 				to.Ptr("0001-arm64"),
	// 			},
	// 			UniqueOfferID: to.Ptr("mktp3pp.kuku-buku"),
	// 		},
	// 	},
	// }
}
