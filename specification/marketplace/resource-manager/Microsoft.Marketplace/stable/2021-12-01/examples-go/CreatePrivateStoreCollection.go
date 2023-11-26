package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/CreatePrivateStoreCollection.json
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
	res, err := clientFactory.NewPrivateStoreCollectionClient().CreateOrUpdate(ctx, "a0e28e55-90c4-41d8-8e34-bb7ef7775406", "d0f5aa2c-ecc3-4d87-906a-f8c486dcc4f1", &armmarketplace.PrivateStoreCollectionClientCreateOrUpdateOptions{Payload: &armmarketplace.Collection{
		Properties: &armmarketplace.CollectionProperties{
			AllSubscriptions: to.Ptr(false),
			Claim:            to.Ptr(""),
			CollectionName:   to.Ptr("Test Collection"),
			SubscriptionsList: []*string{
				to.Ptr("b340914e-353d-453a-85fb-8f9b65b51f91"),
				to.Ptr("f2baa04d-5bfc-461b-b6d8-61b403c9ec48")},
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Collection = armmarketplace.Collection{
	// 	Name: to.Ptr("d0f5aa2c-ecc3-4d87-906a-f8c486dcc4f1"),
	// 	Type: to.Ptr("Microsoft.Marketplace/privateStores/collections"),
	// 	ID: to.Ptr("providers/Microsoft.Marketplace/privateStores/a0e28e55-90c4-41d8-8e34-bb7ef7775406/collections/d0f5aa2c-ecc3-4d87-906a-f8c486dcc4f1"),
	// 	SystemData: &armmarketplace.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T08:23:17.657Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@somedoamin.com"),
	// 		CreatedByType: to.Ptr(armmarketplace.IdentityTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T08:23:17.657Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@somedoamin.com"),
	// 		LastModifiedByType: to.Ptr(armmarketplace.IdentityTypeUser),
	// 	},
	// 	Properties: &armmarketplace.CollectionProperties{
	// 		AllSubscriptions: to.Ptr(false),
	// 		Claim: to.Ptr(""),
	// 		CollectionName: to.Ptr("Test Collection"),
	// 		SubscriptionsList: []*string{
	// 			to.Ptr("b340914e-353d-453a-85fb-8f9b65b51f91"),
	// 			to.Ptr("f2baa04d-5bfc-461b-b6d8-61b403c9ec48")},
	// 		},
	// 	}
}
