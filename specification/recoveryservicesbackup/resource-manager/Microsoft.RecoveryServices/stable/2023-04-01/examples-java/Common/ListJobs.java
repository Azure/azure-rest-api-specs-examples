/** Samples for BackupJobs List. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/Common/ListJobs.json
     */
    /**
     * Sample code: List All Jobs.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void listAllJobs(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.backupJobs().list("NetSDKTestRsVault", "SwaggerTestRg", null, null, com.azure.core.util.Context.NONE);
    }
}
