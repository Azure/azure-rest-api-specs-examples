
/**
 * Samples for RecoveryPoints Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/AzureWorkload/RecoveryPoints_Get_Snapshot.json
     */
    /**
     * Sample code: Get Azure Workload SQL snapshot recovery point details.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getAzureWorkloadSQLSnapshotRecoveryPointDetails(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.recoveryPoints().getWithResponse("testVault", "testRG", "Azure", "VMAppContainer;Compute;testRG;sqlVm",
            "SQLDataBase;mssqlserver;inventory", "1700000000000", com.azure.core.util.Context.NONE);
    }
}
