package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: 2026-08-01-preview/LongTermRetentionBackupListByDatabase.json
func ExampleLongTermRetentionBackupsClient_NewListByDatabasePager_getAllLongTermRetentionBackupsUnderTheDatabase() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLongTermRetentionBackupsClient().NewListByDatabasePager("japaneast", "testserver", "testDatabase", nil)
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
		// page = armsql.LongTermRetentionBackupsClientListByDatabaseResponse{
		// 	LongTermRetentionBackupListResult: armsql.LongTermRetentionBackupListResult{
		// 		Value: []*armsql.LongTermRetentionBackup{
		// 			{
		// 				Name: to.Ptr("55555555-6666-7777-8888-999999999999;131637960820000000;Hot"),
		// 				Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionServers/longTermRetentionDatabases/longTermRetentionBackups"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/locations/japaneast/longTermRetentionServers/testserver/longTermRetentionDatabases/testDatabase/longTermRetentionBackups/55555555-6666-7777-8888-999999999999;131637960820000000;Hot"),
		// 				Properties: &armsql.LongTermRetentionBackupProperties{
		// 					BackupStorageAccessTier: to.Ptr(armsql.BackupStorageAccessTierHot),
		// 					BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 					BackupTime: to.Ptr(time.Date(2017, time.August, 23, 8, 0, 0, 0, time.UTC)),
		// 					DatabaseName: to.Ptr("testDatabase"),
		// 					IsBackupImmutable: to.Ptr(true),
		// 					LegalHoldImmutability: to.Ptr(armsql.SetLegalHoldImmutabilityEnabled),
		// 					ServerCreateTime: to.Ptr(time.Date(2017, time.March, 10, 8, 0, 0, 0, time.UTC)),
		// 					ServerName: to.Ptr("testserver"),
		// 					TimeBasedImmutability: to.Ptr(armsql.TimeBasedImmutabilityDisabled),
		// 					TimeBasedImmutabilityMode: to.Ptr(armsql.TimeBasedImmutabilityModeUnlocked),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("55555555-6666-7777-8888-999999999999;131677960820000000;Hot"),
		// 				Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionServers/longTermRetentionDatabases/longTermRetentionBackups"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/locations/japaneast/longTermRetentionServers/testserver/longTermRetentionDatabases/testDatabase/longTermRetentionBackups/55555555-6666-7777-8888-999999999999;131677960820000000;Hot"),
		// 				Properties: &armsql.LongTermRetentionBackupProperties{
		// 					BackupStorageAccessTier: to.Ptr(armsql.BackupStorageAccessTierHot),
		// 					BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 					BackupTime: to.Ptr(time.Date(2017, time.September, 6, 8, 0, 0, 0, time.UTC)),
		// 					DatabaseDeletionTime: to.Ptr(time.Date(2017, time.September, 7, 8, 0, 0, 0, time.UTC)),
		// 					DatabaseName: to.Ptr("testDatabase"),
		// 					IsBackupImmutable: to.Ptr(true),
		// 					LegalHoldImmutability: to.Ptr(armsql.SetLegalHoldImmutabilityDisabled),
		// 					ServerCreateTime: to.Ptr(time.Date(2017, time.March, 10, 8, 0, 0, 0, time.UTC)),
		// 					ServerName: to.Ptr("testserver"),
		// 					TimeBasedImmutability: to.Ptr(armsql.TimeBasedImmutabilityEnabled),
		// 					TimeBasedImmutabilityMode: to.Ptr(armsql.TimeBasedImmutabilityModeUnlocked),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
