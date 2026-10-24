
/**
 * Samples for LogCaptureConfigurations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/LogCaptureConfigurations_Get.json
     */
    /**
     * Sample code: Get the PostgreSQL server log capture configuration.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        getThePostgreSQLServerLogCaptureConfiguration(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.logCaptureConfigurations().getWithResponse("example-resource-group", "example-cluster",
            "PostgreSqlServer", com.azure.core.util.Context.NONE);
    }
}
