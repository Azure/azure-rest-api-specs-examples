package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/GetPrivateStoreCollectionOffer.json
func ExamplePrivateStoreCollectionOfferClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmarketplace.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateStoreCollectionOfferClient().Get(ctx, "a0e28e55-90c4-41d8-8e34-bb7ef7775406", "marketplacetestthirdparty.md-test-third-party-2", "56a1a02d-8cf8-45df-bf37-d5f7120fcb3d", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Offer = armmarketplace.Offer{
	// 	Name: to.Ptr("marketplacetestthirdparty.md-test-third-party-2"),
	// 	Type: to.Ptr("Microsoft.Marketplace/privateStores/collections/offers"),
	// 	ID: to.Ptr("/providers/Microsoft.Marketplace/privateStores/a0e28e55-90c4-41d8-8e34-bb7ef7775406/collections/56a1a02d-8cf8-45df-bf37-d5f7120fcb3d/offers/marketplacetestthirdparty.md-test-third-party-2"),
	// 	Properties: &armmarketplace.OfferProperties{
	// 		CreatedAt: to.Ptr("05/28/2015 5:50"),
	// 		ETag: to.Ptr("\"9301f4fd-0000-0100-0000-5e248b350000\""),
	// 		ModifiedAt: to.Ptr("05/29/2015 5:50"),
	// 		OfferDisplayName: to.Ptr("md-test-third-party-2"),
	// 		PrivateStoreID: to.Ptr("a0e28e55-90c4-41d8-8e34-bb7ef7775406"),
	// 		PublisherDisplayName: to.Ptr("Marketplace Test Third Party"),
	// 		SpecificPlanIDsLimitation: []*string{
	// 			to.Ptr("0001"),
	// 			to.Ptr("0002"),
	// 			to.Ptr("0003")},
	// 			UniqueOfferID: to.Ptr("marketplacetestthirdparty.md-test-third-party-2"),
	// 		},
	// 	}
}
