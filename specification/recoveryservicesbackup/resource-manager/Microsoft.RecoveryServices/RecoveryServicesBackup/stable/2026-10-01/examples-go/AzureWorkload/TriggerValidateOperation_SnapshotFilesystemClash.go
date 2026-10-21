package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v5"
)

// Generated from example definition: 2026-10-01/AzureWorkload/TriggerValidateOperation_SnapshotFilesystemClash.json
func ExampleValidateOperationClient_BeginTrigger_validateSqlSnapshotRestoreAfterResolvingFilesystemClashes() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewValidateOperationClient().BeginTrigger(ctx, "testVault", "testRG", armrecoveryservicesbackup.ValidateOperationRequestResource{
		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.RecoveryServices/vaults/testVault/backupFabrics/Azure/protectionContainers/VMAppContainer;Compute;testRG;sqlVm/protectedItems/SQLDataBase;mssqlserver;inventory/recoveryPoints/1700000000000"),
		Properties: &armrecoveryservicesbackup.ValidateAzureWorkloadRestoreOperationRequest{
			ObjectType: to.Ptr("ValidateAzureWorkloadRestoreOperationRequest"),
			RestoreRequest: &armrecoveryservicesbackup.AzureWorkloadSQLRestoreRequest{
				ObjectType:   to.Ptr("AzureWorkloadSQLRestoreRequest"),
				RecoveryMode: to.Ptr(armrecoveryservicesbackup.RecoveryMode("Snapshot")),
				RecoveryType: to.Ptr(armrecoveryservicesbackup.RecoveryTypeAlternateLocation),
				SnapshotRestoreParameters: &armrecoveryservicesbackup.SnapshotRestoreParameters{
					DisksToDetachOnClash: []*string{
						to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/targetRG/providers/Microsoft.Compute/disks/sqlDataDisk01"),
						to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/targetRG/providers/Microsoft.Compute/disks/sqlLogDisk01"),
					},
					SkipAttachAndMount: to.Ptr(false),
				},
				ShouldUseAlternateTargetLocation: to.Ptr(true),
				SourceResourceID:                 to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.Compute/virtualMachines/sqlVm"),
				TargetResourceGroupName:          to.Ptr("targetRG"),
				TargetVirtualMachineID:           to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/targetRG/providers/Microsoft.Compute/virtualMachines/sqlRestoreVm"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
}
