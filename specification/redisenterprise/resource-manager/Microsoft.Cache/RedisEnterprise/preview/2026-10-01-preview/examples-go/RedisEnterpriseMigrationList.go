package armredisenterprise_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise/v4"
)

// Generated from example definition: 2026-10-01-preview/RedisEnterpriseMigrationList.json
func ExampleMigrationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredisenterprise.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMigrationsClient().NewListPager("rg1", "cache1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armredisenterprise.MigrationsClientListResponse{
		// 	MigrationList: armredisenterprise.MigrationList{
		// 		Value: []*armredisenterprise.Migration{
		// 			{
		// 				Name: to.Ptr("cache1/default"),
		// 				Type: to.Ptr("Microsoft.Cache/redisEnterprise/migrations"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Cache/redisEnterprise/cache1/migrations/default"),
		// 				Properties: &armredisenterprise.AzureCacheForRedisMigrationProperties{
		// 					CreationTime: to.Ptr(time.Date(2024, time.October, 1, 12, 0, 0, 0, time.UTC)),
		// 					LastModifiedTime: to.Ptr(time.Date(2024, time.October, 1, 12, 0, 0, 0, time.UTC)),
		// 					ProvisioningState: to.Ptr(armredisenterprise.MigrationProvisioningStateInProgress),
		// 					SkipDataMigration: to.Ptr(true),
		// 					SourceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Cache/redis/cache1"),
		// 					SourceType: to.Ptr(armredisenterprise.SourceTypeAzureCacheForRedis),
		// 					StatusDetails: nil,
		// 					SwitchDNS: to.Ptr(true),
		// 					TargetResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Cache/redisEnterprise/cache1"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
