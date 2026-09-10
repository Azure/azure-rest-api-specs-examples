
/**
 * Samples for Monitors RefreshIngestionKey.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/Monitors_RefreshIngestionKey.json
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
