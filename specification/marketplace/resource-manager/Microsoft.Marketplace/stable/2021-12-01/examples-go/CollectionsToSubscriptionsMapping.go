package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/CollectionsToSubscriptionsMapping.json
func ExamplePrivateStoreClient_CollectionsToSubscriptionsMapping() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmarketplace.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateStoreClient().CollectionsToSubscriptionsMapping(ctx, "a0e28e55-90c4-41d8-8e34-bb7ef7775406", &armmarketplace.PrivateStoreClientCollectionsToSubscriptionsMappingOptions{Payload: &armmarketplace.CollectionsToSubscriptionsMappingPayload{
		Properties: &armmarketplace.CollectionsToSubscriptionsMappingProperties{
			SubscriptionIDs: []*string{
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
	// res.CollectionsToSubscriptionsMappingResponse = armmarketplace.CollectionsToSubscriptionsMappingResponse{
	// 	Details: map[string]*armmarketplace.CollectionsSubscriptionsMappingDetails{
	// 		"4eb49758-f591-486f-bd58-dff00fb7a8d8": &armmarketplace.CollectionsSubscriptionsMappingDetails{
	// 			CollectionName: to.Ptr("Test Collection"),
	// 			Subscriptions: []*string{
	// 				to.Ptr("b340914e-353d-453a-85fb-8f9b65b51f91"),
	// 				to.Ptr("f2baa04d-5bfc-461b-b6d8-61b403c9ec48")},
	// 			},
	// 			"74c02e27-2524-436c-831d-d64565f31153": &armmarketplace.CollectionsSubscriptionsMappingDetails{
	// 				CollectionName: to.Ptr("Test Collection 2"),
	// 			},
	// 		},
	// 	}
}
