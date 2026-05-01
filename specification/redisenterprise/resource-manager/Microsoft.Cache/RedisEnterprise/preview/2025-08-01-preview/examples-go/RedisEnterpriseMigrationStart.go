package armredisenterprise_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise/v4"
)

// Generated from example definition: 2025-08-01-preview/RedisEnterpriseMigrationStart.json
func ExampleMigrationClient_BeginStart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredisenterprise.NewClientFactory("e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMigrationClient().BeginStart(ctx, "rg1", "cache1", armredisenterprise.Migration{
		Properties: &armredisenterprise.AzureCacheForRedisMigrationProperties{
			SkipDataMigration: to.Ptr(true),
			SourceResourceID:  to.Ptr("/subscriptions/e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f/resourceGroups/rg1/providers/Microsoft.Cache/redis/cache1"),
			SourceType:        to.Ptr(armredisenterprise.SourceTypeAzureCacheForRedis),
			SwitchDNS:         to.Ptr(true),
		},
	}, nil)
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
	// res = armredisenterprise.MigrationClientStartResponse{
	// 	Migration: &armredisenterprise.Migration{
	// 		Name: to.Ptr("default"),
	// 		Type: to.Ptr("Microsoft.Cache/redisEnterprise/migrations"),
	// 		ID: to.Ptr("/subscriptions/e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f/resourceGroups/rg1/providers/Microsoft.Cache/redisEnterprise/cache1/migrations/default"),
	// 		Properties: &armredisenterprise.AzureCacheForRedisMigrationProperties{
	// 			CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-01T12:00:00Z"); return t}()),
	// 			LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-01T12:00:00Z"); return t}()),
	// 			ProvisioningState: to.Ptr(armredisenterprise.MigrationProvisioningState("Completed")),
	// 			SkipDataMigration: to.Ptr(true),
	// 			SourceResourceID: to.Ptr("/subscriptions/e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f/resourceGroups/rg1/providers/Microsoft.Cache/redis/cache1"),
	// 			SourceType: to.Ptr(armredisenterprise.SourceTypeAzureCacheForRedis),
	// 			StatusDetails: to.Ptr(""),
	// 			SwitchDNS: to.Ptr(true),
	// 			TargetResourceID: to.Ptr("/subscriptions/e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f/resourceGroups/rg1/providers/Microsoft.Cache/redisEnterprise/cache1"),
	// 		},
	// 	},
	// }
}
