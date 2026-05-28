package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: 2025-02-01-preview/ResourceGroupBasedManagedInstanceLongTermRetentionBackupListByLocationMax.json
func ExampleLongTermRetentionManagedInstanceBackupsClient_NewListByResourceGroupLocationPager_getAllLongTermRetentionBackupsUnderTheLocationWithMaximalParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLongTermRetentionManagedInstanceBackupsClient().NewListByResourceGroupLocationPager("testResourceGroup", "japaneast", &armsql.LongTermRetentionManagedInstanceBackupsClientListByResourceGroupLocationOptions{
		Filter: to.Ptr("Properties/ManagedInstanceName eq 'testInstance1'"),
		Skip:   to.Ptr[int64](0),
		Top:    to.Ptr[int64](2)})
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
		// page = armsql.LongTermRetentionManagedInstanceBackupsClientListByResourceGroupLocationResponse{
		// 	ManagedInstanceLongTermRetentionBackupListResult: armsql.ManagedInstanceLongTermRetentionBackupListResult{
		// 		Value: []*armsql.ManagedInstanceLongTermRetentionBackup{
		// 			{
		// 				Name: to.Ptr("55555555-6666-7777-8888-999999999999;2017-08-23T08:00:00.000Z;Archive"),
		// 				Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionManagedInstances/longTermRetentionDatabases/longTermRetentionManagedInstanceBackups"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testResourceGroup/providers/Microsoft.Sql/locations/japaneast/longTermRetentionManagedInstances/testInstance1/longTermRetentionDatabases/testDatabase1/longTermRetentionManagedInstanceBackups/55555555-6666-7777-8888-999999999999;131637960820000000;Archive"),
		// 				Properties: &armsql.ManagedInstanceLongTermRetentionBackupProperties{
		// 					BackupStorageAccessTier: to.Ptr(armsql.BackupStorageAccessTierArchive),
		// 					BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 					BackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-23T08:00:00Z"); return t}()),
		// 					DatabaseName: to.Ptr("testDatabase1"),
		// 					ManagedInstanceCreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-10T08:00:00Z"); return t}()),
		// 					ManagedInstanceName: to.Ptr("testInstance1"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("12341234-1234-1234-1234-123123123123;2017-08-30T08:00:00.000Z;Hot"),
		// 				Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionManagedInstances/longTermRetentionDatabases/longTermRetentionManagedInstanceBackups"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testResourceGroup/providers/Microsoft.Sql/locations/japaneast/longTermRetentionManagedInstances/testInstance1/longTermRetentionDatabases/testDatabase2/longTermRetentionManagedInstanceBackups/12341234-1234-1234-1234-123123123123;131657960820000000;Hot"),
		// 				Properties: &armsql.ManagedInstanceLongTermRetentionBackupProperties{
		// 					BackupStorageAccessTier: to.Ptr(armsql.BackupStorageAccessTierHot),
		// 					BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 					BackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-30T08:00:00Z"); return t}()),
		// 					DatabaseName: to.Ptr("testDatabase2"),
		// 					ManagedInstanceCreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-10T08:00:00Z"); return t}()),
		// 					ManagedInstanceName: to.Ptr("testInstance1"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
