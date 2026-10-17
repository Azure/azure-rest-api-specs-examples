const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to validate operation for specified backed up item in the form of an asynchronous operation. Returns tracking headers which can be tracked using GetValidateOperationResult API.
 *
 * @summary validate operation for specified backed up item in the form of an asynchronous operation. Returns tracking headers which can be tracked using GetValidateOperationResult API.
 * x-ms-original-file: 2026-10-01/AzureWorkload/TriggerValidateOperation_SnapshotFilesystemClash.json
 */
async function validateSQLSnapshotRestoreAfterResolvingFilesystemClashes() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  await client.validateOperation.trigger("testVault", "testRG", {
    id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.RecoveryServices/vaults/testVault/backupFabrics/Azure/protectionContainers/VMAppContainer;Compute;testRG;sqlVm/protectedItems/SQLDataBase;mssqlserver;inventory/recoveryPoints/1700000000000",
    properties: {
      objectType: "ValidateAzureWorkloadRestoreOperationRequest",
      restoreRequest: {
        objectType: "AzureWorkloadSQLRestoreRequest",
        recoveryMode: "Snapshot",
        recoveryType: "AlternateLocation",
        snapshotRestoreParameters: {
          disksToDetachOnClash: [
            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/targetRG/providers/Microsoft.Compute/disks/sqlDataDisk01",
            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/targetRG/providers/Microsoft.Compute/disks/sqlLogDisk01",
          ],
          skipAttachAndMount: false,
        },
        shouldUseAlternateTargetLocation: true,
        sourceResourceId:
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.Compute/virtualMachines/sqlVm",
        targetResourceGroupName: "targetRG",
        targetVirtualMachineId:
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/targetRG/providers/Microsoft.Compute/virtualMachines/sqlRestoreVm",
      },
    },
  });
}
