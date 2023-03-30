package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/BulkCollectionsAction.json
func ExamplePrivateStoreClient_BulkCollectionsAction() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmarketplace.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateStoreClient().BulkCollectionsAction(ctx, "a0e28e55-90c4-41d8-8e34-bb7ef7775406", &armmarketplace.PrivateStoreClientBulkCollectionsActionOptions{Payload: &armmarketplace.BulkCollectionsPayload{
		Properties: &armmarketplace.BulkCollectionsDetails{
			Action: to.Ptr("EnableCollections"),
			CollectionIDs: []*string{
				to.Ptr("c752f021-1c37-4af5-b82f-74c51c27b44a"),
				to.Ptr("f47ef1c7-e908-4f39-ae29-db181634ad8d")},
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BulkCollectionsResponse = armmarketplace.BulkCollectionsResponse{
	// 	Failed: []*armmarketplace.CollectionsDetails{
	// 		{
	// 			CollectionID: to.Ptr("f47ef1c7-e908-4f39-ae29-db181634ad8d"),
	// 			CollectionName: to.Ptr("Test collection 2"),
	// 	}},
	// 	Succeeded: []*armmarketplace.CollectionsDetails{
	// 		{
	// 			CollectionID: to.Ptr("c752f021-1c37-4af5-b82f-74c51c27b44a"),
	// 			CollectionName: to.Ptr("Test collection"),
	// 	}},
	// }
}
