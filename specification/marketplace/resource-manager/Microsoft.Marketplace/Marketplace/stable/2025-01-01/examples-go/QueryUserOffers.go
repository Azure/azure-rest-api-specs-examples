package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace/v2"
)

// Generated from example definition: 2025-01-01/QueryUserOffers.json
func ExamplePrivateStoreClient_QueryUserOffers() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmarketplace.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateStoreClient().QueryUserOffers(ctx, "a0e28e55-90c4-41d8-8e34-bb7ef7775406", &armmarketplace.PrivateStoreClientQueryUserOffersOptions{
		Payload: &armmarketplace.QueryUserOffersProperties{
			Properties: &armmarketplace.QueryUserOffersDetails{
				OfferIDs: []*string{
					to.Ptr("contoso.logger"),
					to.Ptr("contoso.monitor"),
				},
				SubscriptionIDs: []*string{
					to.Ptr("b340914e-353d-453a-85fb-8f9b65b51f91"),
				},
			},
		}})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmarketplace.PrivateStoreClientQueryUserOffersResponse{
	// 	QueryOffers: &armmarketplace.QueryOffers{
	// 		NextLink: to.Ptr(""),
	// 		Value: []*armmarketplace.OfferProperties{
	// 			{
	// 				CreatedAt: to.Ptr("2022-03-05T13:55:33.351Z"),
	// 				ETag: to.Ptr("\"9301f4fd-0000-0100-0000-5e248b350000\""),
	// 				IconFileUris: map[string]*string{
	// 					"small": to.Ptr("https://some-images.contoso.com/image/apps.12345678-7654"),
	// 				},
	// 				ModifiedAt: to.Ptr("2021-04-22T14:17:41.8520203Z"),
	// 				OfferDisplayName: to.Ptr("Contoso Log4All"),
	// 				Plans: []*armmarketplace.Plan{
	// 					{
	// 						Accessibility: to.Ptr(armmarketplace.AccessibilityPublic),
	// 						PlanDisplayName: to.Ptr("Log4J"),
	// 						PlanID: to.Ptr("log4j"),
	// 						SKUID: to.Ptr("001"),
	// 						StackType: to.Ptr("arm"),
	// 					},
	// 				},
	// 				PrivateStoreID: to.Ptr("a0e28e55-90c4-41d8-8e34-bb7ef7775406"),
	// 				PublisherDisplayName: to.Ptr("Contoso"),
	// 				SpecificPlanIDsLimitation: []*string{
	// 					to.Ptr("0002"),
	// 				},
	// 				UniqueOfferID: to.Ptr("contoso.logger"),
	// 				UpdateSuppressedDueIdempotence: to.Ptr(true),
	// 			},
	// 		},
	// 	},
	// }
}
