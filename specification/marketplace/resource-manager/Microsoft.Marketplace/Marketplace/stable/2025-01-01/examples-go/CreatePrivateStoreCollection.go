package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace/v2"
)

// Generated from example definition: 2025-01-01/CreatePrivateStoreCollection.json
func ExamplePrivateStoreCollectionClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmarketplace.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateStoreCollectionClient().CreateOrUpdate(ctx, "a0e28e55-90c4-41d8-8e34-bb7ef7775406", "d0f5aa2c-ecc3-4d87-906a-f8c486dcc4f1", armmarketplace.Collection{
		Properties: &armmarketplace.CollectionProperties{
			AllSubscriptions: to.Ptr(false),
			Claim:            to.Ptr(""),
			CollectionName:   to.Ptr("Test Collection"),
			SubscriptionsList: []*string{
				to.Ptr("b340914e-353d-453a-85fb-8f9b65b51f91"),
				to.Ptr("f2baa04d-5bfc-461b-b6d8-61b403c9ec48"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmarketplace.PrivateStoreCollectionClientCreateOrUpdateResponse{
	// 	Collection: &armmarketplace.Collection{
	// 		Name: to.Ptr("d0f5aa2c-ecc3-4d87-906a-f8c486dcc4f1"),
	// 		Type: to.Ptr("Microsoft.Marketplace/privateStores/collections"),
	// 		ID: to.Ptr("providers/Microsoft.Marketplace/privateStores/a0e28e55-90c4-41d8-8e34-bb7ef7775406/collections/d0f5aa2c-ecc3-4d87-906a-f8c486dcc4f1"),
	// 		Properties: &armmarketplace.CollectionProperties{
	// 			AllSubscriptions: to.Ptr(false),
	// 			ApproveAllItems: to.Ptr(false),
	// 			Claim: to.Ptr(""),
	// 			CollectionName: to.Ptr("Test Collection"),
	// 			Enabled: to.Ptr(true),
	// 			NumberOfOffers: to.Ptr[int64](0),
	// 			SubscriptionsList: []*string{
	// 				to.Ptr("b340914e-353d-453a-85fb-8f9b65b51f91"),
	// 				to.Ptr("f2baa04d-5bfc-461b-b6d8-61b403c9ec48"),
	// 			},
	// 		},
	// 		SystemData: &armmarketplace.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T10:23:17.6571572+02:00"); return t}()),
	// 			CreatedBy: to.Ptr("user@somedoamin.com"),
	// 			CreatedByType: to.Ptr(armmarketplace.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T10:23:17.6571572+02:00"); return t}()),
	// 			LastModifiedBy: to.Ptr("user@somedoamin.com"),
	// 			LastModifiedByType: to.Ptr(armmarketplace.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
