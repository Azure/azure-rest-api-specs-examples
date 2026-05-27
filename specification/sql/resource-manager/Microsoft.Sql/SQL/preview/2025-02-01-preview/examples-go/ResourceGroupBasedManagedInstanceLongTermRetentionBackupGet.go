package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: 2025-02-01-preview/ResourceGroupBasedManagedInstanceLongTermRetentionBackupGet.json
func ExampleLongTermRetentionManagedInstanceBackupsClient_GetByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLongTermRetentionManagedInstanceBackupsClient().GetByResourceGroup(ctx, "testResourceGroup", "japaneast", "testInstance", "testDatabase", "55555555-6666-7777-8888-999999999999;131637960820000000;Archive", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsql.LongTermRetentionManagedInstanceBackupsClientGetByResourceGroupResponse{
	// 	ManagedInstanceLongTermRetentionBackup: armsql.ManagedInstanceLongTermRetentionBackup{
	// 		Name: to.Ptr("55555555-6666-7777-8888-999999999999;131637960820000000;Archive"),
	// 		Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionManagedInstances/longTermRetentionDatabases/longTermRetentionManagedInstanceBackups"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testResourceGroup/providers/Microsoft.Sql/Locations/japaneast/longTermRetentionManagedInstances/testInstance/longTermRetentionDatabases/testDatabase/longTermRetentionManagedInstanceBackups/55555555-6666-7777-8888-999999999999;131637960820000000;Archive"),
	// 		Properties: &armsql.ManagedInstanceLongTermRetentionBackupProperties{
	// 			BackupStorageAccessTier: to.Ptr(armsql.BackupStorageAccessTierArchive),
	// 			BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
	// 			BackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-06T08:00:00Z"); return t}()),
	// 			DatabaseName: to.Ptr("testDatabase"),
	// 			ManagedInstanceCreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-10T08:00:00Z"); return t}()),
	// 			ManagedInstanceName: to.Ptr("testInstance"),
	// 		},
	// 	},
	// }
}
