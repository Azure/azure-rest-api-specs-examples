
/**
 * Samples for Jobs Export.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Common/TriggerExportJobs.json
     */
    /**
     * Sample code: Export Jobs.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void
        exportJobs(com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.jobs().exportWithResponse("NetSDKTestRsVault", "SwaggerTestRg", null, com.azure.core.util.Context.NONE);
    }
}
