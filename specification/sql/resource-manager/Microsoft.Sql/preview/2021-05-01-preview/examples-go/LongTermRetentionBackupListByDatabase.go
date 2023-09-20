package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/LongTermRetentionBackupListByDatabase.json
func ExampleLongTermRetentionBackupsClient_NewListByDatabasePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLongTermRetentionBackupsClient().NewListByDatabasePager("japaneast", "testserver", "testDatabase", &armsql.LongTermRetentionBackupsClientListByDatabaseOptions{OnlyLatestPerDatabase: nil,
		DatabaseState: nil,
	})
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
		// page.LongTermRetentionBackupListResult = armsql.LongTermRetentionBackupListResult{
		// 	Value: []*armsql.LongTermRetentionBackup{
		// 		{
		// 			Name: to.Ptr("55555555-6666-7777-8888-999999999999;131637960820000000"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionServers/longTermRetentionDatabases/longTermRetentionBackups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/locations/japaneast/longTermRetentionServers/testserver/longTermRetentionDatabases/testDatabase/longTermRetentionBackups/55555555-6666-7777-8888-999999999999;131637960820000000"),
		// 			Properties: &armsql.LongTermRetentionBackupProperties{
		// 				BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 				BackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-23T08:00:00Z"); return t}()),
		// 				DatabaseName: to.Ptr("testDatabase"),
		// 				ServerCreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-10T08:00:00Z"); return t}()),
		// 				ServerName: to.Ptr("testserver"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("55555555-6666-7777-8888-999999999999;131657960820000000"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionServers/longTermRetentionDatabases/longTermRetentionBackups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/locations/japaneast/longTermRetentionServers/testserver/longTermRetentionDatabases/testDatabase/longTermRetentionBackups/55555555-6666-7777-8888-999999999999;131657960820000000"),
		// 			Properties: &armsql.LongTermRetentionBackupProperties{
		// 				BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 				BackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-30T08:00:00Z"); return t}()),
		// 				DatabaseName: to.Ptr("testDatabase"),
		// 				ServerCreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-10T08:00:00Z"); return t}()),
		// 				ServerName: to.Ptr("testserver"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("55555555-6666-7777-8888-999999999999;131677960820000000"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionServers/longTermRetentionDatabases/longTermRetentionBackups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/locations/japaneast/longTermRetentionServers/testserver/longTermRetentionDatabases/testDatabase/longTermRetentionBackups/55555555-6666-7777-8888-999999999999;131677960820000000"),
		// 			Properties: &armsql.LongTermRetentionBackupProperties{
		// 				BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 				BackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-06T08:00:00Z"); return t}()),
		// 				DatabaseDeletionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-07T08:00:00Z"); return t}()),
		// 				DatabaseName: to.Ptr("testDatabase"),
		// 				ServerCreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-10T08:00:00Z"); return t}()),
		// 				ServerName: to.Ptr("testserver"),
		// 			},
		// 	}},
		// }
	}
}
