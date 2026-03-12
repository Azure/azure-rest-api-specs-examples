package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace/v2"
)

// Generated from example definition: 2025-01-01/DisableApproveAllItems.json
func ExamplePrivateStoreCollectionClient_DisableApproveAllItems() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmarketplace.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateStoreCollectionClient().DisableApproveAllItems(ctx, "a0e28e55-90c4-41d8-8e34-bb7ef7775406", "56a1a02d-8cf8-45df-bf37-d5f7120fcb3d", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmarketplace.PrivateStoreCollectionClientDisableApproveAllItemsResponse{
	// 	Collection: &armmarketplace.Collection{
	// 		Name: to.Ptr("56a1a02d-8cf8-45df-bf37-d5f7120fcb3d"),
	// 		Type: to.Ptr("Microsoft.Marketplace/privateStores/collections"),
	// 		ID: to.Ptr("providers/Microsoft.Marketplace/privateStores/a0e28e55-90c4-41d8-8e34-bb7ef7775406/collections/56a1a02d-8cf8-45df-bf37-d5f7120fcb3d"),
	// 		Properties: &armmarketplace.CollectionProperties{
	// 			AllSubscriptions: to.Ptr(false),
	// 			AppliedRules: []*armmarketplace.Rule{
	// 				{
	// 					Type: to.Ptr(armmarketplace.RuleTypePrivateProducts),
	// 				},
	// 			},
	// 			ApproveAllItems: to.Ptr(false),
	// 			ApproveAllItemsModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-07T14:30:58.566Z"); return t}()),
	// 			Claim: to.Ptr(""),
	// 			CollectionID: to.Ptr("56a1a02d-8cf8-45df-bf37-d5f7120fcb3d"),
	// 			CollectionName: to.Ptr("Global"),
	// 			Enabled: to.Ptr(true),
	// 			NumberOfOffers: to.Ptr[int64](0),
	// 			SubscriptionsList: []*string{
	// 				to.Ptr("57a5024d-2b80-4434-a37a-445e74a032bb"),
	// 				to.Ptr("c4b603f7-4449-4b98-8077-571e17a661c1"),
	// 			},
	// 		},
	// 		SystemData: &armmarketplace.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-07T14:00:05.566Z"); return t}()),
	// 			CreatedBy: to.Ptr("user@somedoamin.com"),
	// 			CreatedByType: to.Ptr(armmarketplace.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-07T14:30:58.566Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user@somedoamin.com"),
	// 			LastModifiedByType: to.Ptr(armmarketplace.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
