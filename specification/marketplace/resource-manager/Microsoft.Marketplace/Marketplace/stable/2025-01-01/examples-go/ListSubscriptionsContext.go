package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace/v2"
)

// Generated from example definition: 2025-01-01/ListSubscriptionsContext.json
func ExamplePrivateStoreClient_ListSubscriptionsContext() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmarketplace.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateStoreClient().ListSubscriptionsContext(ctx, "a0e28e55-90c4-41d8-8e34-bb7ef7775406", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmarketplace.PrivateStoreClientListSubscriptionsContextResponse{
	// 	SubscriptionsContextList: &armmarketplace.SubscriptionsContextList{
	// 		SubscriptionsIDs: []*string{
	// 			to.Ptr("090dcec8-3415-4e13-9377-07f489cdfeed"),
	// 			to.Ptr("030dcec2-4323-5e13-9376-96f489cdfeed"),
	// 		},
	// 	},
	// }
}
