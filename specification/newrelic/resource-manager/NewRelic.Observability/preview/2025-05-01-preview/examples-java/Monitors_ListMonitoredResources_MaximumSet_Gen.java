
/**
 * Samples for Monitors ListMonitoredResources.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/newrelic/resource-manager/NewRelic.Observability/preview/2025-05-01-preview/examples/
     * Monitors_ListMonitoredResources_MaximumSet_Gen.json
     */
    /**
     * Sample code: Monitors_ListMonitoredResources_MaximumSet_Gen.
     * 
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void monitorsListMonitoredResourcesMaximumSetGen(
        com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager.monitors().listMonitoredResources("rgopenapi", "ipxmlcbonyxtolzejcjshkmlron",
            com.azure.core.util.Context.NONE);
    }
}
