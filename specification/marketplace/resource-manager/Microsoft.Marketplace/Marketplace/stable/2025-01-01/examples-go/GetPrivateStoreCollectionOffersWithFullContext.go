package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace/v2"
)

// Generated from example definition: 2025-01-01/GetPrivateStoreCollectionOffersWithFullContext.json
func ExamplePrivateStoreCollectionOfferClient_NewListByContextsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmarketplace.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateStoreCollectionOfferClient().NewListByContextsPager("a0e28e55-90c4-41d8-8e34-bb7ef7775406", "56a1a02d-8cf8-45df-bf37-d5f7120fcb3d", &armmarketplace.PrivateStoreCollectionOfferClientListByContextsOptions{
		Payload: &armmarketplace.CollectionOffersByAllContextsPayload{
			Properties: &armmarketplace.CollectionOffersByAllContextsProperties{
				SubscriptionIDs: []*string{
					to.Ptr("b340914e-353d-453a-85fb-8f9b65b51f91"),
					to.Ptr("f2baa04d-5bfc-461b-b6d8-61b403c9ec48"),
				},
			},
		}})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armmarketplace.PrivateStoreCollectionOfferClientListByContextsResponse{
		// 	CollectionOffersByContextList: armmarketplace.CollectionOffersByContextList{
		// 		NextLink: to.Ptr(""),
		// 		Value: []*armmarketplace.CollectionOffersByContext{
		// 			{
		// 				Context: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				Offers: &armmarketplace.CollectionOffersByContextOffers{
		// 					Value: []*armmarketplace.OfferProperties{
		// 						{
		// 							CreatedAt: to.Ptr("05/28/2015 5:50"),
		// 							ETag: to.Ptr("\"9301f4fd-0000-0100-0304-5e248b350043\""),
		// 							ModifiedAt: to.Ptr("05/29/2015 5:50"),
		// 							OfferDisplayName: to.Ptr("md-test-third-party-3"),
		// 							PrivateStoreID: to.Ptr("a0e28e55-90c4-41d8-8e34-bb7ef7775406"),
		// 							PublisherDisplayName: to.Ptr("Marketplace Test Third Party"),
		// 							SpecificPlanIDsLimitation: []*string{
		// 								to.Ptr("public1"),
		// 								to.Ptr("public2"),
		// 							},
		// 							UniqueOfferID: to.Ptr("marketplacetestthirdparty.md-test-third-party-3"),
		// 						},
		// 						{
		// 							CreatedAt: to.Ptr("05/28/2015 5:50"),
		// 							ModifiedAt: to.Ptr("05/29/2015 5:50"),
		// 							OfferDisplayName: to.Ptr("md-test-third-party-2"),
		// 							PrivateStoreID: to.Ptr("a0e28e55-90c4-41d8-8e34-bb7ef7775406"),
		// 							PublisherDisplayName: to.Ptr("Marketplace Test Third Party"),
		// 							SpecificPlanIDsLimitation: []*string{
		// 								to.Ptr("mdTestPublicPlan"),
		// 							},
		// 							UniqueOfferID: to.Ptr("marketplacetestthirdparty.md-test-third-party-2"),
		// 						},
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Context: to.Ptr("a0e28e55-90c4-41d8-8e34-bb7ef7775406"),
		// 				Offers: &armmarketplace.CollectionOffersByContextOffers{
		// 					Value: []*armmarketplace.OfferProperties{
		// 						{
		// 							CreatedAt: to.Ptr("05/28/2015 5:50"),
		// 							ModifiedAt: to.Ptr("05/29/2015 5:50"),
		// 							OfferDisplayName: to.Ptr("md-test-third-party-2"),
		// 							PrivateStoreID: to.Ptr("a0e28e55-90c4-41d8-8e34-bb7ef7775406"),
		// 							PublisherDisplayName: to.Ptr("Marketplace Test Third Party"),
		// 							SpecificPlanIDsLimitation: []*string{
		// 								to.Ptr("mdTestPrivatePlanForTenant"),
		// 							},
		// 							UniqueOfferID: to.Ptr("marketplacetestthirdparty.md-test-third-party-2"),
		// 						},
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Context: to.Ptr("b340914e-353d-453a-85fb-8f9b65b51f91"),
		// 				Offers: &armmarketplace.CollectionOffersByContextOffers{
		// 					Value: []*armmarketplace.OfferProperties{
		// 						{
		// 							CreatedAt: to.Ptr("05/28/2015 5:50"),
		// 							ModifiedAt: to.Ptr("05/29/2015 5:50"),
		// 							OfferDisplayName: to.Ptr("md-test-third-party-2"),
		// 							PrivateStoreID: to.Ptr("a0e28e55-90c4-41d8-8e34-bb7ef7775406"),
		// 							PublisherDisplayName: to.Ptr("Marketplace Test Third Party"),
		// 							SpecificPlanIDsLimitation: []*string{
		// 								to.Ptr("mdTestPrivatePlanSub1"),
		// 							},
		// 							UniqueOfferID: to.Ptr("marketplacetestthirdparty.md-test-third-party-2"),
		// 						},
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Context: to.Ptr("f2baa04d-5bfc-461b-b6d8-61b403c9ec48"),
		// 				Offers: &armmarketplace.CollectionOffersByContextOffers{
		// 					Value: []*armmarketplace.OfferProperties{
		// 						{
		// 							CreatedAt: to.Ptr("05/28/2015 5:50"),
		// 							ModifiedAt: to.Ptr("05/29/2015 5:50"),
		// 							OfferDisplayName: to.Ptr("md-test-third-party-2"),
		// 							PrivateStoreID: to.Ptr("a0e28e55-90c4-41d8-8e34-bb7ef7775406"),
		// 							PublisherDisplayName: to.Ptr("Marketplace Test Third Party"),
		// 							SpecificPlanIDsLimitation: []*string{
		// 								to.Ptr("mdTestPrivatePlanSub2"),
		// 							},
		// 							UniqueOfferID: to.Ptr("marketplacetestthirdparty.md-test-third-party-2"),
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
