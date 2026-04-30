
/**
 * Samples for BackupUsageSummaries List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-31-preview/Common/BackupProtectionContainers_UsageSummary_Get.json
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
