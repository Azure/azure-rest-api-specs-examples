
import com.azure.resourcemanager.recoveryservicesbackup.models.AzureWorkloadSqlRestoreRequest;
import com.azure.resourcemanager.recoveryservicesbackup.models.RecoveryMode;
import com.azure.resourcemanager.recoveryservicesbackup.models.RecoveryType;
import com.azure.resourcemanager.recoveryservicesbackup.models.SnapshotRestoreParameters;
import com.azure.resourcemanager.recoveryservicesbackup.models.ValidateAzureWorkloadRestoreOperationRequest;
import com.azure.resourcemanager.recoveryservicesbackup.models.ValidateOperationRequestResource;
import java.util.Arrays;

/**
 * Samples for ValidateOperation Trigger.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/AzureWorkload/TriggerValidateOperation_SnapshotFilesystemClash.json
     */
    /**
     * Sample code: Validate SQL snapshot restore after resolving filesystem clashes.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void validateSQLSnapshotRestoreAfterResolvingFilesystemClashes(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.validateOperations().trigger("testVault", "testRG", new ValidateOperationRequestResource().withId(
            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.RecoveryServices/vaults/testVault/backupFabrics/Azure/protectionContainers/VMAppContainer;Compute;testRG;sqlVm/protectedItems/SQLDataBase;mssqlserver;inventory/recoveryPoints/1700000000000")
            .withProperties(new ValidateAzureWorkloadRestoreOperationRequest().withRestoreRequest(
                new AzureWorkloadSqlRestoreRequest().withRecoveryType(RecoveryType.ALTERNATE_LOCATION)
                    .withSourceResourceId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.Compute/virtualMachines/sqlVm")
                    .withRecoveryMode(RecoveryMode.fromString("Snapshot")).withTargetResourceGroupName("targetRG")
                    .withSnapshotRestoreParameters(new SnapshotRestoreParameters().withSkipAttachAndMount(false)
                        .withDisksToDetachOnClash(Arrays.asList(
                            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/targetRG/providers/Microsoft.Compute/disks/sqlDataDisk01",
                            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/targetRG/providers/Microsoft.Compute/disks/sqlLogDisk01")))
                    .withTargetVirtualMachineId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/targetRG/providers/Microsoft.Compute/virtualMachines/sqlRestoreVm")
                    .withShouldUseAlternateTargetLocation(true))),
            com.azure.core.util.Context.NONE);
    }
}
