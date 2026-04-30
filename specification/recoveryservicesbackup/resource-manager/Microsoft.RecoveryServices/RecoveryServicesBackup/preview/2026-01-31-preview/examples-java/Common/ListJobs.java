
/**
 * Samples for BackupJobs List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-31-preview/Common/ListJobs.json
     */
    /**
     * Sample code: List All Jobs.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void
        listAllJobs(com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.backupJobs().list("NetSDKTestRsVault", "SwaggerTestRg", null, null, com.azure.core.util.Context.NONE);
    }
}
