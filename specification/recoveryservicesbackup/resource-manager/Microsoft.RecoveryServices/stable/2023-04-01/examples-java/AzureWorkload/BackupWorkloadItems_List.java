/** Samples for BackupWorkloadItems List. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/AzureWorkload/BackupWorkloadItems_List.json
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
                com.azure.core.util.Context.NONE);
    }
}
