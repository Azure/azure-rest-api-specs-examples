import com.azure.core.util.Context;

/** Samples for Jobs Export. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/preview/2022-09-01-preview/examples/Common/TriggerExportJobs.json
     */
    /**
     * Sample code: Export Jobs.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void exportJobs(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.jobs().exportWithResponse("NetSDKTestRsVault", "SwaggerTestRg", null, Context.NONE);
    }
}
