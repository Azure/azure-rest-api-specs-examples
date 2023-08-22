/** Samples for Monitors ListMonitoredResources. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2023-04-27/examples/Monitors_ListMonitoredResources_MaximumSet_Gen.json
     */
    /**
     * Sample code: Monitors_ListMonitoredResources_MaximumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void monitorsListMonitoredResourcesMaximumSetGen(
        com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager.monitors().listMonitoredResources("myResourceGroup", "myMonitor", com.azure.core.util.Context.NONE);
    }
}
