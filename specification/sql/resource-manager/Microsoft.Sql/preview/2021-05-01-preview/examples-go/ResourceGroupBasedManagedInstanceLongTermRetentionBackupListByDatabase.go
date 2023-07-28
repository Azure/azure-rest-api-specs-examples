package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/ResourceGroupBasedManagedInstanceLongTermRetentionBackupListByDatabase.json
func ExampleLongTermRetentionManagedInstanceBackupsClient_NewListByResourceGroupDatabasePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLongTermRetentionManagedInstanceBackupsClient().NewListByResourceGroupDatabasePager("testResourceGroup", "japaneast", "testInstance", "testDatabase", &armsql.LongTermRetentionManagedInstanceBackupsClientListByResourceGroupDatabaseOptions{OnlyLatestPerDatabase: nil,
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
		// page.ManagedInstanceLongTermRetentionBackupListResult = armsql.ManagedInstanceLongTermRetentionBackupListResult{
		// 	Value: []*armsql.ManagedInstanceLongTermRetentionBackup{
		// 		{
		// 			Name: to.Ptr("2018-06-01T08:00:00.000Z;55555555-6666-7777-8888-999999999999;2018-08-23T08:00:00.000Z"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionManagedInstances/longTermRetentionDatabases/longTermRetentionManagedInstanceBackups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testResourceGroup/providers/Microsoft.Sql/Locations/japaneast/longTermRetentionManagedInstances/testInstance/longTermRetentionDatabases/testDatabase/longTermRetentionManagedInstanceBackups/2018-06-01T08:00:00.000Z;55555555-6666-7777-8888-999999999999;2018-08-23T08:00:00.000Z"),
		// 			Properties: &armsql.ManagedInstanceLongTermRetentionBackupProperties{
		// 				BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 				BackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-23T08:00:00Z"); return t}()),
		// 				DatabaseName: to.Ptr("testDatabase"),
		// 				ManagedInstanceCreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-10T08:00:00Z"); return t}()),
		// 				ManagedInstanceName: to.Ptr("testInstance"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("2018-06-01T08:00:00.000Z;55555555-6666-7777-8888-999999999999;2018-08-30T08:00:00.000Z"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionManagedInstances/longTermRetentionDatabases/longTermRetentionManagedInstanceBackups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testResourceGroup/providers/Microsoft.Sql/Locations/japaneast/longTermRetentionManagedInstances/testInstance/longTermRetentionDatabases/testDatabase/longTermRetentionManagedInstanceBackups/2018-06-01T08:00:00.000Z;55555555-6666-7777-8888-999999999999;2018-08-30T08:00:00.000Z"),
		// 			Properties: &armsql.ManagedInstanceLongTermRetentionBackupProperties{
		// 				BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 				BackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-30T08:00:00Z"); return t}()),
		// 				DatabaseName: to.Ptr("testDatabase"),
		// 				ManagedInstanceCreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-10T08:00:00Z"); return t}()),
		// 				ManagedInstanceName: to.Ptr("testInstance"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("2018-06-01T08:00:00.000Z;55555555-6666-7777-8888-999999999999;2018-09-06T08:00:00.000Z"),
		// 			Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionManagedInstances/longTermRetentionDatabases/longTermRetentionManagedInstanceBackups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testResourceGroup/providers/Microsoft.Sql/Locations/japaneast/longTermRetentionManagedInstances/testInstance/longTermRetentionDatabases/testDatabase/longTermRetentionManagedInstanceBackups/2018-06-01T08:00:00.000Z;55555555-6666-7777-8888-999999999999;2018-09-06T08:00:00.000Z"),
		// 			Properties: &armsql.ManagedInstanceLongTermRetentionBackupProperties{
		// 				BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 				BackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-06T08:00:00Z"); return t}()),
		// 				DatabaseDeletionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-07T08:00:00Z"); return t}()),
		// 				DatabaseName: to.Ptr("testDatabase"),
		// 				ManagedInstanceCreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-10T08:00:00Z"); return t}()),
		// 				ManagedInstanceName: to.Ptr("testInstance"),
		// 			},
		// 	}},
		// }
	}
}
