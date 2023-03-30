/** Samples for Monitors ListMonitoredResources. */
public final class Main {
    /*
     * x-ms-original-file: specification/newrelic/resource-manager/NewRelic.Observability/preview/2022-07-01-preview/examples/Monitors_ListMonitoredResources_MinimumSet_Gen.json
     */
    /**
     * Sample code: Monitors_ListMonitoredResources_MinimumSet_Gen.
     *
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void monitorsListMonitoredResourcesMinimumSetGen(
        com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager
            .monitors()
            .listMonitoredResources("rgopenapi", "ipxmlcbonyxtolzejcjshkmlron", com.azure.core.util.Context.NONE);
    }
}
