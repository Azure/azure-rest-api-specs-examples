package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/GetPrivateStoreCollectionsList.json
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
	// res.CollectionsList = armmarketplace.CollectionsList{
	// 	Value: []*armmarketplace.Collection{
	// 		{
	// 			Name: to.Ptr("56a1a02d-8cf8-45df-bf37-d5f7120fcb3d"),
	// 			Type: to.Ptr("Microsoft.Marketplace/privateStores/collections"),
	// 			ID: to.Ptr("providers/Microsoft.Marketplace/privateStores/a0e28e55-90c4-41d8-8e34-bb7ef7775406/collections"),
	// 			SystemData: &armmarketplace.SystemData{
	// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-02-01T08:23:17.657Z"); return t}()),
	// 				CreatedBy: to.Ptr("user@somedoamin.com"),
	// 				CreatedByType: to.Ptr(armmarketplace.IdentityTypeUser),
	// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-02-01T08:23:17.657Z"); return t}()),
	// 				LastModifiedBy: to.Ptr("user@somedoamin.com"),
	// 				LastModifiedByType: to.Ptr(armmarketplace.IdentityTypeUser),
	// 			},
	// 			Properties: &armmarketplace.CollectionProperties{
	// 				AllSubscriptions: to.Ptr(true),
	// 				Claim: to.Ptr(""),
	// 				CollectionID: to.Ptr("56a1a02d-8cf8-45df-bf37-d5f7120fcb3d"),
	// 				CollectionName: to.Ptr("Default Collection"),
	// 				SubscriptionsList: []*string{
	// 				},
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("fba3f52c-874a-4010-87cf-c1cfa6ed3490"),
	// 			Type: to.Ptr("Microsoft.Marketplace/privateStores/collections"),
	// 			ID: to.Ptr("providers/Microsoft.Marketplace/privateStores/a0e28e55-90c4-41d8-8e34-bb7ef7775406/collections"),
	// 			SystemData: &armmarketplace.SystemData{
	// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-02-01T08:23:17.657Z"); return t}()),
	// 				CreatedBy: to.Ptr("user@somedoamin.com"),
	// 				CreatedByType: to.Ptr(armmarketplace.IdentityTypeUser),
	// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-02-01T08:23:17.657Z"); return t}()),
	// 				LastModifiedBy: to.Ptr("user@somedoamin.com"),
	// 				LastModifiedByType: to.Ptr(armmarketplace.IdentityTypeUser),
	// 			},
	// 			Properties: &armmarketplace.CollectionProperties{
	// 				AllSubscriptions: to.Ptr(false),
	// 				Claim: to.Ptr(""),
	// 				CollectionID: to.Ptr("56a1a02d-8cf8-45df-bf37-d5f7120fcb3d"),
	// 				CollectionName: to.Ptr("Dev collection"),
	// 				SubscriptionsList: []*string{
	// 					to.Ptr("7c927b63-59cf-4a0f-9d13-41e11f1ddf76")},
	// 				},
	// 		}},
	// 	}
}
