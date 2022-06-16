import com.azure.core.util.Context;

/** Samples for BackupUsageSummaries List. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-07-01/examples/Common/BackupProtectedItem_UsageSummary_Get.json
     */
    /**
     * Sample code: Get Protected Items Usages Summary.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getProtectedItemsUsagesSummary(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .backupUsageSummaries()
            .list("testVault", "testRG", "type eq 'BackupProtectedItemCountSummary'", null, Context.NONE);
    }
}
