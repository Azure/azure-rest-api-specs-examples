
/**
 * Samples for LogCaptureConfigurations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/LogCaptureConfigurations_List.json
     */
    /**
     * Sample code: List log capture configurations.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void listLogCaptureConfigurations(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.logCaptureConfigurations().list("example-resource-group", "example-cluster",
            com.azure.core.util.Context.NONE);
    }
}
