/** Samples for ExportJobs Trigger. */
public final class Main {
    /*
     * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/JobCRUD/TriggerExportJobs.json
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
