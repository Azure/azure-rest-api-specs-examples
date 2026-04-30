
/**
 * Samples for ExportJobsOperationResults Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-31-preview/Common/ExportJobsOperationResult.json
     */
    /**
     * Sample code: Export Jobs Operation Results.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void exportJobsOperationResults(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.exportJobsOperationResults().getWithResponse("NetSDKTestRsVault", "SwaggerTestRg",
            "00000000-0000-0000-0000-000000000000", com.azure.core.util.Context.NONE);
    }
}
