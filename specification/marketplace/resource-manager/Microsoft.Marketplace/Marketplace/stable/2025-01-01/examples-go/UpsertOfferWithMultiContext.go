package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace/v2"
)

// Generated from example definition: 2025-01-01/UpsertOfferWithMultiContext.json
func ExamplePrivateStoreCollectionOfferClient_UpsertOfferWithMultiContext() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmarketplace.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateStoreCollectionOfferClient().UpsertOfferWithMultiContext(ctx, "a0e28e55-90c4-41d8-8e34-bb7ef7775406", "56a1a02d-8cf8-45df-bf37-d5f7120fcb3d", "contoso.logger", &armmarketplace.PrivateStoreCollectionOfferClientUpsertOfferWithMultiContextOptions{
		Payload: &armmarketplace.MultiContextAndPlansPayload{
			Properties: &armmarketplace.MultiContextAndPlansProperties{
				ETag:    to.Ptr("\"9301f4fd-0000-0100-0000-5e248b350332\""),
				OfferID: to.Ptr("contoso.logger"),
				PlansContext: []*armmarketplace.ContextAndPlansDetails{
					{
						Context: to.Ptr("a5edbe7d-9f73-47fd-834a-0d6142f4c7a1"),
						PlanIDs: []*string{
							to.Ptr("log4db"),
							to.Ptr("log4file"),
						},
					},
					{
						Context: to.Ptr("45b604af-19bb-448e-a761-4a6be7374b2f"),
						PlanIDs: []*string{
							to.Ptr("log4web"),
						},
					},
				},
			},
		}})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmarketplace.PrivateStoreCollectionOfferClientUpsertOfferWithMultiContextResponse{
	// 	Offer: &armmarketplace.Offer{
	// 		Name: to.Ptr("contoso.logger"),
	// 		Type: to.Ptr("Microsoft.Marketplace/privateStores/collections/offers"),
	// 		ID: to.Ptr("/providers/Microsoft.Marketplace/privateStores/a0e28e55-90c4-41d8-8e34-bb7ef7775406/collections/56a1a02d-8cf8-45df-bf37-d5f7120fcb3d/offers/contoso.logger"),
	// 		Properties: &armmarketplace.OfferProperties{
	// 			CreatedAt: to.Ptr("2022-03-05T13:55:33.351Z"),
	// 			ETag: to.Ptr("\"9301f4fd-0000-0100-0000-5e248b350000\""),
	// 			IconFileUris: map[string]*string{
	// 				"small": to.Ptr("https://some-images.contoso.com/image/apps.12345678-7654"),
	// 			},
	// 			ModifiedAt: to.Ptr("2021-04-22T14:17:41.8520203Z"),
	// 			OfferDisplayName: to.Ptr("Contoso Log4All"),
	// 			Plans: []*armmarketplace.Plan{
	// 				{
	// 					Accessibility: to.Ptr(armmarketplace.AccessibilityPublic),
	// 					PlanDisplayName: to.Ptr("log4db"),
	// 					PlanID: to.Ptr("log4db"),
	// 					SKUID: to.Ptr("001"),
	// 					StackType: to.Ptr("arm"),
	// 				},
	// 				{
	// 					Accessibility: to.Ptr(armmarketplace.AccessibilityPublic),
	// 					PlanDisplayName: to.Ptr("log4file"),
	// 					PlanID: to.Ptr("log4file"),
	// 					SKUID: to.Ptr("002"),
	// 					StackType: to.Ptr("arm"),
	// 				},
	// 				{
	// 					Accessibility: to.Ptr(armmarketplace.AccessibilityPublic),
	// 					PlanDisplayName: to.Ptr("log4web"),
	// 					PlanID: to.Ptr("log4web"),
	// 					SKUID: to.Ptr("003"),
	// 					StackType: to.Ptr("arm"),
	// 				},
	// 			},
	// 			PrivateStoreID: to.Ptr("a0e28e55-90c4-41d8-8e34-bb7ef7775406"),
	// 			PublisherDisplayName: to.Ptr("Contoso"),
	// 			SpecificPlanIDsLimitation: []*string{
	// 				to.Ptr("log4db"),
	// 				to.Ptr("log4file"),
	// 				to.Ptr("log4web"),
	// 			},
	// 			UniqueOfferID: to.Ptr("contoso.logger"),
	// 			UpdateSuppressedDueIdempotence: to.Ptr(true),
	// 		},
	// 	},
	// }
}
