package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v4"
)

// Generated from example definition: 2026-06-01/StorageContextCacheCRUD/ContextCaches_CreateOrUpdate.json
func ExampleContextCachesClient_BeginCreateOrUpdate_createAContextCache() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewContextCachesClient().BeginCreateOrUpdate(ctx, "testrg", "testaccount", armstorage.ContextCache{
		Location: to.Ptr("eastus"),
		Tags: map[string]*string{
			"environment": to.Ptr("test"),
		},
		Properties: &armstorage.ContextCacheProperties{
			AccountKind: to.Ptr(armstorage.ContextCacheAccountKindRegional),
			Description: to.Ptr("Test Azure Context Cache account"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armstorage.ContextCachesClientCreateOrUpdateResponse{
	// 	ContextCache: armstorage.ContextCache{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Storage/contextCaches/testaccount"),
	// 		Name: to.Ptr("testaccount"),
	// 		Type: to.Ptr("Microsoft.Storage/contextCaches"),
	// 		Location: to.Ptr("eastus"),
	// 		Tags: map[string]*string{
	// 			"environment": to.Ptr("test"),
	// 		},
	// 		Properties: &armstorage.ContextCacheProperties{
	// 			AccountKind: to.Ptr(armstorage.ContextCacheAccountKindRegional),
	// 			Description: to.Ptr("Test Azure Context Cache account"),
	// 			ProvisioningState: to.Ptr(armstorage.ContextCacheProvisioningStateSucceeded),
	// 		},
	// 		SystemData: &armstorage.SystemData{
	// 			CreatedBy: to.Ptr("user@example.com"),
	// 			CreatedByType: to.Ptr(armstorage.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC)),
	// 			LastModifiedBy: to.Ptr("user@example.com"),
	// 			LastModifiedByType: to.Ptr(armstorage.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC)),
	// 		},
	// 	},
	// }
}
