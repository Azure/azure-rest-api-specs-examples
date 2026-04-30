
/**
 * Samples for BackupUsageSummaries List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-31-preview/Common/BackupProtectedItem_UsageSummary_Get.json
     */
    /**
     * Sample code: Get Protected Items Usages Summary.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getProtectedItemsUsagesSummary(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.backupUsageSummaries().list("testVault", "testRG", "type eq 'BackupProtectedItemCountSummary'", null,
            com.azure.core.util.Context.NONE);
    }
}
