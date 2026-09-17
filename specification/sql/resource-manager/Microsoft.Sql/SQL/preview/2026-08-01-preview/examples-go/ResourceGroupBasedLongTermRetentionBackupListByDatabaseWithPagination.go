package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: 2026-08-01-preview/ResourceGroupBasedLongTermRetentionBackupListByDatabaseWithPagination.json
func ExampleLongTermRetentionBackupsClient_NewListByResourceGroupDatabasePager_getLongTermRetentionBackupsUnderTheDatabaseBasedOnResourceGroupWithPagination() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLongTermRetentionBackupsClient().NewListByResourceGroupDatabasePager("testResourceGroup", "japaneast", "testserver", "testDatabase", nil)
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
		// page = armsql.LongTermRetentionBackupsClientListByResourceGroupDatabaseResponse{
		// 	LongTermRetentionBackupListResult: armsql.LongTermRetentionBackupListResult{
		// 		Value: []*armsql.LongTermRetentionBackup{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testResourceGroup/providers/Microsoft.Sql/locations/japaneast/longTermRetentionServers/testserver/longTermRetentionDatabases/testDatabase/longTermRetentionBackups/55555555-6666-7777-8888-999999999999;131637960820000000;Hot"),
		// 				Name: to.Ptr("55555555-6666-7777-8888-999999999999;131637960820000000;Hot"),
		// 				Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionServers/longTermRetentionDatabases/longTermRetentionBackups"),
		// 				Properties: &armsql.LongTermRetentionBackupProperties{
		// 					ServerName: to.Ptr("testserver"),
		// 					ServerCreateTime: to.Ptr(time.Date(2017, time.March, 10, 8, 0, 0, 0, time.UTC)),
		// 					DatabaseName: to.Ptr("testDatabase"),
		// 					BackupTime: to.Ptr(time.Date(2017, time.August, 23, 8, 0, 0, 0, time.UTC)),
		// 					BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 					IsBackupImmutable: to.Ptr(true),
		// 					BackupStorageAccessTier: to.Ptr(armsql.BackupStorageAccessTierHot),
		// 					TimeBasedImmutability: to.Ptr(armsql.TimeBasedImmutabilityEnabled),
		// 					TimeBasedImmutabilityMode: to.Ptr(armsql.TimeBasedImmutabilityModeLocked),
		// 					LegalHoldImmutability: to.Ptr(armsql.SetLegalHoldImmutabilityDisabled),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testResourceGroup/providers/Microsoft.Sql/locations/japaneast/longTermRetentionServers/testserver/longTermRetentionDatabases/testDatabase/longTermRetentionBackups/55555555-6666-7777-8888-999999999999;131657960820000000;Hot"),
		// 				Name: to.Ptr("55555555-6666-7777-8888-999999999999;131657960820000000;Hot"),
		// 				Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionServers/longTermRetentionDatabases/longTermRetentionBackups"),
		// 				Properties: &armsql.LongTermRetentionBackupProperties{
		// 					ServerName: to.Ptr("testserver"),
		// 					ServerCreateTime: to.Ptr(time.Date(2017, time.March, 10, 8, 0, 0, 0, time.UTC)),
		// 					DatabaseName: to.Ptr("testDatabase"),
		// 					BackupTime: to.Ptr(time.Date(2017, time.August, 30, 8, 0, 0, 0, time.UTC)),
		// 					BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 					IsBackupImmutable: to.Ptr(true),
		// 					BackupStorageAccessTier: to.Ptr(armsql.BackupStorageAccessTierHot),
		// 					TimeBasedImmutability: to.Ptr(armsql.TimeBasedImmutabilityEnabled),
		// 					TimeBasedImmutabilityMode: to.Ptr(armsql.TimeBasedImmutabilityModeLocked),
		// 					LegalHoldImmutability: to.Ptr(armsql.SetLegalHoldImmutabilityDisabled),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testResourceGroup/providers/Microsoft.Sql/locations/japaneast/longTermRetentionServers/testserver/longTermRetentionDatabases/testDatabase/longTermRetentionBackups?api-version=2026-08-01-preview&$top=2&$skipToken=eyJEYXRhYmFzZUlkIjoiNTU1NTU1NTUtNjY2Ni03Nzc3LTg4ODgtOTk5OTk5OTk5OTk5IiwiQmFja3VwVGltZSI6IjIwMTctMDgtMzBUMDg6MDA6MDBaIiwiQmFja3VwU3RvcmFnZUFjY2Vzc1RpZXIiOiJIb3QifQ%3D%3D"),
		// 	},
		// }
	}
}
