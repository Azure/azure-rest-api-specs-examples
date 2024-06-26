import com.azure.core.util.Context;

/** Samples for BackupJobs List. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/preview/2022-09-01-preview/examples/Common/ListJobs.json
     */
    /**
     * Sample code: List All Jobs.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void listAllJobs(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.backupJobs().list("NetSDKTestRsVault", "SwaggerTestRg", null, null, Context.NONE);
    }
}
