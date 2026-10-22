package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v5"
)

// Generated from example definition: 2026-10-01/AzureWorkload/RecoveryPoints_Get_Snapshot.json
func ExampleRecoveryPointsClient_Get_getAzureWorkloadSqlSnapshotRecoveryPointDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRecoveryPointsClient().Get(ctx, "testVault", "testRG", "Azure", "VMAppContainer;Compute;testRG;sqlVm", "SQLDataBase;mssqlserver;inventory", "1700000000000", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armrecoveryservicesbackup.RecoveryPointsClientGetResponse{
	// 	RecoveryPointResource: armrecoveryservicesbackup.RecoveryPointResource{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.RecoveryServices/vaults/testVault/backupFabrics/Azure/protectionContainers/VMAppContainer;Compute;testRG;sqlVm/protectedItems/SQLDataBase;mssqlserver;inventory/recoveryPoints/1700000000000"),
	// 		Name: to.Ptr("1700000000000"),
	// 		Properties: &armrecoveryservicesbackup.AzureWorkloadSQLRecoveryPoint{
	// 			ExtendedInfo: &armrecoveryservicesbackup.AzureWorkloadSQLRecoveryPointExtendedInfo{
	// 				SnapshotRecoveryPointInfo: &armrecoveryservicesbackup.SnapshotRecoveryPointInfo{
	// 					SnapshotResourceGroup: to.Ptr("AzureBackupSnapshots_testRG"),
	// 					SourceFilesystemInfo: []*armrecoveryservicesbackup.FilesystemInfo{
	// 						{
	// 							AccessPaths: []*string{
	// 								to.Ptr("F:\\"),
	// 							},
	// 							DiskInfoList: []*armrecoveryservicesbackup.DiskInfo{
	// 								{
	// 									DiskGUID: to.Ptr("6f9619ff-8b86-d011-b42d-00c04fc964ff"),
	// 									DiskLocation: to.Ptr("PCIROOT(0)#PCI(1D00)#PCI(0000)"),
	// 									DiskNumber: to.Ptr[int32](2),
	// 									DiskUniqueID: to.Ptr("60022480-9abc-def0-1234-567890abcdef"),
	// 									FriendlyName: to.Ptr("SQL data disk"),
	// 									LUN: to.Ptr[int32](1),
	// 									ManagedDiskID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.Compute/disks/sqlDataDisk01"),
	// 									SizeInBytes: to.Ptr[int64](1099511627776),
	// 								},
	// 							},
	// 							FileSystemType: to.Ptr("NTFS"),
	// 							IsOnStorageSpace: to.Ptr(false),
	// 							Label: to.Ptr("SQLData"),
	// 							VolumeGUID: to.Ptr("\\\\?\\Volume{12345678-1234-1234-1234-123456789abc}\\"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ObjectType: to.Ptr("AzureWorkloadSQLRecoveryPoint"),
	// 			RecoveryPointTimeInUTC: to.Ptr(time.Date(2026, time.September, 15, 2, 30, 0, 0, time.UTC)),
	// 			Type: to.Ptr(armrecoveryservicesbackup.RestorePointTypeFull),
	// 		},
	// 		Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupFabrics/protectionContainers/protectedItems/recoveryPoints"),
	// 	},
	// }
}
