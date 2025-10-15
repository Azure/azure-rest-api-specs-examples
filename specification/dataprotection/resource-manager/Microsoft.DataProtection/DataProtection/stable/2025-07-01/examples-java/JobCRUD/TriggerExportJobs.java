
/**
 * Samples for ExportJobs Trigger.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/JobCRUD/TriggerExportJobs.json
     */
    /**
     * Sample code: Trigger Export Jobs.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void triggerExportJobs(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.exportJobs().trigger("SwaggerTestRg", "NetSDKTestRsVault", com.azure.core.util.Context.NONE);
    }
}
