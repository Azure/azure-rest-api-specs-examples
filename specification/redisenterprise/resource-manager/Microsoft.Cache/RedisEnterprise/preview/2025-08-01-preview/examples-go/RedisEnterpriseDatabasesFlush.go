package armredisenterprise_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise/v4"
)

// Generated from example definition: 2025-08-01-preview/RedisEnterpriseDatabasesFlush.json
func ExampleDatabasesClient_BeginFlush() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredisenterprise.NewClientFactory("e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDatabasesClient().BeginFlush(ctx, "rg1", "cache1", "default", &armredisenterprise.DatabasesClientBeginFlushOptions{
		Parameters: &armredisenterprise.FlushParameters{
			IDs: []*string{
				to.Ptr("/subscriptions/e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f2/resourceGroups/rg2/providers/Microsoft.Cache/redisEnterprise/cache2/databases/default"),
			},
		}})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armredisenterprise.DatabasesClientFlushResponse{
	// }
}
