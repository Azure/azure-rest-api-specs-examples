
/**
 * Samples for Monitors ListMonitoredResources.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01-preview/Monitors_ListMonitoredResources_MinimumSet_Gen.json
     */
    /**
     * Sample code: Monitors_ListMonitoredResources_MinimumSet_Gen.
     * 
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void monitorsListMonitoredResourcesMinimumSetGen(
        com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager.monitors().listMonitoredResources("rgopenapi", "ipxmlcbonyxtolzejcjshkmlron",
            com.azure.core.util.Context.NONE);
    }
}
