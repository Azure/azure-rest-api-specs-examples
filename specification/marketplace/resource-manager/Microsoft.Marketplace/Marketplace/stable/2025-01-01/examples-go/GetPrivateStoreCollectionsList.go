package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace/v2"
)

// Generated from example definition: 2025-01-01/GetPrivateStoreCollectionsList.json
func ExamplePrivateStoreCollectionClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmarketplace.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateStoreCollectionClient().List(ctx, "a0e28e55-90c4-41d8-8e34-bb7ef7775406", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmarketplace.PrivateStoreCollectionClientListResponse{
	// 	CollectionsList: &armmarketplace.CollectionsList{
	// 		NextLink: to.Ptr(""),
	// 		Value: []*armmarketplace.Collection{
	// 			{
	// 				Name: to.Ptr("56a1a02d-8cf8-45df-bf37-d5f7120fcb3d"),
	// 				Type: to.Ptr("Microsoft.Marketplace/privateStores/collections"),
	// 				ID: to.Ptr("providers/Microsoft.Marketplace/privateStores/a0e28e55-90c4-41d8-8e34-bb7ef7775406/collections"),
	// 				Properties: &armmarketplace.CollectionProperties{
	// 					AllSubscriptions: to.Ptr(true),
	// 					AppliedRules: []*armmarketplace.Rule{
	// 						{
	// 							Type: to.Ptr(armmarketplace.RuleTypePrivateProducts),
	// 						},
	// 					},
	// 					ApproveAllItems: to.Ptr(false),
	// 					Claim: to.Ptr(""),
	// 					CollectionID: to.Ptr("56a1a02d-8cf8-45df-bf37-d5f7120fcb3d"),
	// 					CollectionName: to.Ptr("Default Collection"),
	// 					Enabled: to.Ptr(true),
	// 					NumberOfOffers: to.Ptr[int64](2),
	// 					SubscriptionsList: []*string{
	// 					},
	// 				},
	// 				SystemData: &armmarketplace.SystemData{
	// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-02-01T10:23:17.6571572+02:00"); return t}()),
	// 					CreatedBy: to.Ptr("user@somedoamin.com"),
	// 					CreatedByType: to.Ptr(armmarketplace.CreatedByTypeUser),
	// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-02-01T10:23:17.6571572+02:00"); return t}()),
	// 					LastModifiedBy: to.Ptr("user@somedoamin.com"),
	// 					LastModifiedByType: to.Ptr(armmarketplace.CreatedByTypeUser),
	// 				},
	// 			},
	// 			{
	// 				Name: to.Ptr("fba3f52c-874a-4010-87cf-c1cfa6ed3490"),
	// 				Type: to.Ptr("Microsoft.Marketplace/privateStores/collections"),
	// 				ID: to.Ptr("providers/Microsoft.Marketplace/privateStores/a0e28e55-90c4-41d8-8e34-bb7ef7775406/collections"),
	// 				Properties: &armmarketplace.CollectionProperties{
	// 					AllSubscriptions: to.Ptr(false),
	// 					AppliedRules: []*armmarketplace.Rule{
	// 					},
	// 					ApproveAllItems: to.Ptr(true),
	// 					ApproveAllItemsModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-25T07:53:01.9685921Z"); return t}()),
	// 					Claim: to.Ptr(""),
	// 					CollectionID: to.Ptr("56a1a02d-8cf8-45df-bf37-d5f7120fcb3d"),
	// 					CollectionName: to.Ptr("Dev collection"),
	// 					Enabled: to.Ptr(true),
	// 					NumberOfOffers: to.Ptr[int64](1),
	// 					SubscriptionsList: []*string{
	// 						to.Ptr("7c927b63-59cf-4a0f-9d13-41e11f1ddf76"),
	// 					},
	// 				},
	// 				SystemData: &armmarketplace.SystemData{
	// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-02-01T10:23:17.6571572+02:00"); return t}()),
	// 					CreatedBy: to.Ptr("user@somedoamin.com"),
	// 					CreatedByType: to.Ptr(armmarketplace.CreatedByTypeUser),
	// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-05-01T10:23:17.6571572+02:00"); return t}()),
	// 					LastModifiedBy: to.Ptr("user@somedoamin.com"),
	// 					LastModifiedByType: to.Ptr(armmarketplace.CreatedByTypeUser),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
