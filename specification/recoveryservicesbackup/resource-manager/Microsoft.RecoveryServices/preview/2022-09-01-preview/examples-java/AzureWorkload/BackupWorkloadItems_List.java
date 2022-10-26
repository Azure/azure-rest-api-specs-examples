import com.azure.core.util.Context;

/** Samples for BackupWorkloadItems List. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/preview/2022-09-01-preview/examples/AzureWorkload/BackupWorkloadItems_List.json
     */
    /**
     * Sample code: List Workload Items in Container.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void listWorkloadItemsInContainer(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .backupWorkloadItems()
            .list(
                "suchandr-seacan-rsv",
                "testRg",
                "Azure",
                "VMAppContainer;Compute;bvtdtestag;sqlserver-1",
                "backupManagementType eq 'AzureWorkload'",
                null,
                Context.NONE);
    }
}
