
/**
 * Samples for Monitors RefreshIngestionKey.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01-preview/Monitors_RefreshIngestionKey.json
     */
    /**
     * Sample code: Monitors_RefreshIngestionKey.
     * 
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void monitorsRefreshIngestionKey(
        com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager.monitors().refreshIngestionKeyWithResponse("myResourceGroup", "myMonitor",
            com.azure.core.util.Context.NONE);
    }
}
