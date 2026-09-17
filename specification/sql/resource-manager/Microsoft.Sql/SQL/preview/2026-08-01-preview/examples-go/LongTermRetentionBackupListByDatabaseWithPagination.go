package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: 2026-08-01-preview/LongTermRetentionBackupListByDatabaseWithPagination.json
func ExampleLongTermRetentionBackupsClient_NewListByDatabasePager_getLongTermRetentionBackupsUnderTheDatabaseWithPagination() {
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
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/locations/japaneast/longTermRetentionServers/testserver/longTermRetentionDatabases/testDatabase/longTermRetentionBackups/55555555-6666-7777-8888-999999999999;131637960820000000;Hot"),
		// 				Name: to.Ptr("55555555-6666-7777-8888-999999999999;131637960820000000;Hot"),
		// 				Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionServers/longTermRetentionDatabases/longTermRetentionBackups"),
		// 				Properties: &armsql.LongTermRetentionBackupProperties{
		// 					ServerName: to.Ptr("testserver"),
		// 					ServerCreateTime: to.Ptr(time.Date(2017, time.March, 10, 8, 0, 0, 0, time.UTC)),
		// 					DatabaseName: to.Ptr("testDatabase"),
		// 					BackupTime: to.Ptr(time.Date(2017, time.August, 23, 8, 0, 0, 0, time.UTC)),
		// 					BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 					IsBackupImmutable: to.Ptr(true),
		// 					TimeBasedImmutability: to.Ptr(armsql.TimeBasedImmutabilityDisabled),
		// 					TimeBasedImmutabilityMode: to.Ptr(armsql.TimeBasedImmutabilityModeUnlocked),
		// 					LegalHoldImmutability: to.Ptr(armsql.SetLegalHoldImmutabilityEnabled),
		// 					BackupStorageAccessTier: to.Ptr(armsql.BackupStorageAccessTierHot),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/locations/japaneast/longTermRetentionServers/testserver/longTermRetentionDatabases/testDatabase/longTermRetentionBackups/55555555-6666-7777-8888-999999999999;131677960820000000;Hot"),
		// 				Name: to.Ptr("55555555-6666-7777-8888-999999999999;131677960820000000;Hot"),
		// 				Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionServers/longTermRetentionDatabases/longTermRetentionBackups"),
		// 				Properties: &armsql.LongTermRetentionBackupProperties{
		// 					ServerName: to.Ptr("testserver"),
		// 					ServerCreateTime: to.Ptr(time.Date(2017, time.March, 10, 8, 0, 0, 0, time.UTC)),
		// 					DatabaseName: to.Ptr("testDatabase"),
		// 					DatabaseDeletionTime: to.Ptr(time.Date(2017, time.September, 7, 8, 0, 0, 0, time.UTC)),
		// 					BackupTime: to.Ptr(time.Date(2017, time.September, 6, 8, 0, 0, 0, time.UTC)),
		// 					BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 					IsBackupImmutable: to.Ptr(true),
		// 					TimeBasedImmutability: to.Ptr(armsql.TimeBasedImmutabilityEnabled),
		// 					TimeBasedImmutabilityMode: to.Ptr(armsql.TimeBasedImmutabilityModeUnlocked),
		// 					LegalHoldImmutability: to.Ptr(armsql.SetLegalHoldImmutabilityDisabled),
		// 					BackupStorageAccessTier: to.Ptr(armsql.BackupStorageAccessTierHot),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/locations/japaneast/longTermRetentionServers/testserver/longTermRetentionDatabases/testDatabase/longTermRetentionBackups?api-version=2026-08-01-preview&$top=2&$skipToken=eyJEYXRhYmFzZUlkIjoiNTU1NTU1NTUtNjY2Ni03Nzc3LTg4ODgtOTk5OTk5OTk5OTk5IiwiQmFja3VwVGltZSI6IjIwMTctMDktMDZUMDg6MDA6MDBaIiwiQmFja3VwU3RvcmFnZUFjY2Vzc1RpZXIiOiJIb3QifQ%3D%3D"),
		// 	},
		// }
	}
}
