package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: 2026-08-01-preview/ResourceGroupBasedRemoveTimeBasedImmutabilityManagedInstanceLongTermRetentionBackup.json
func ExampleLongTermRetentionManagedInstanceBackupsClient_BeginRemoveTimeBasedImmutabilityByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLongTermRetentionManagedInstanceBackupsClient().BeginRemoveTimeBasedImmutabilityByResourceGroup(ctx, "testResourceGroup", "japaneast", "testInstance", "testDatabase", "55555555-6666-7777-8888-999999999999;131637960820000000;Archive", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsql.LongTermRetentionManagedInstanceBackupsClientRemoveTimeBasedImmutabilityByResourceGroupResponse{
	// 	ManagedInstanceLongTermRetentionBackup: armsql.ManagedInstanceLongTermRetentionBackup{
	// 		Name: to.Ptr("55555555-6666-7777-8888-999999999999;131637960820000000;Archive"),
	// 		Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionManagedInstances/longTermRetentionDatabases/longTermRetentionManagedInstanceBackups"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testResourceGroup/providers/Microsoft.Sql/locations/japaneast/longTermRetentionManagedInstances/testInstance/longTermRetentionDatabases/testDatabase/longTermRetentionManagedInstanceBackups/55555555-6666-7777-8888-999999999999;131637960820000000;Archive"),
	// 		Properties: &armsql.ManagedInstanceLongTermRetentionBackupProperties{
	// 			BackupStorageAccessTier: to.Ptr(armsql.BackupStorageAccessTierArchive),
	// 			BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
	// 			BackupTime: to.Ptr(time.Date(2018, time.August, 23, 8, 0, 0, 0, time.UTC)),
	// 			DatabaseName: to.Ptr("testDatabase"),
	// 			ManagedInstanceCreateTime: to.Ptr(time.Date(2017, time.September, 1, 8, 0, 0, 0, time.UTC)),
	// 			ManagedInstanceName: to.Ptr("testInstance"),
	// 			TimeBasedImmutability: to.Ptr(armsql.TimeBasedImmutabilityDisabled),
	// 		},
	// 	},
	// }
}
