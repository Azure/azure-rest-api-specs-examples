
/**
 * Samples for BackupUsageSummaries List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2024-04-01/examples/
     * Common/BackupProtectionContainers_UsageSummary_Get.json
     */
    /**
     * Sample code: Get Protected Containers Usages Summary.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getProtectedContainersUsagesSummary(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.backupUsageSummaries().list("testVault", "testRG", "type eq 'BackupProtectionContainerCountSummary'",
            null, com.azure.core.util.Context.NONE);
    }
}
