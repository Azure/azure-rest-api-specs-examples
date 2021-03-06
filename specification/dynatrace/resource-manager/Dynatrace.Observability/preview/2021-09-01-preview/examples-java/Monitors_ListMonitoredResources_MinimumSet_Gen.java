import com.azure.core.util.Context;

/** Samples for Monitors ListMonitoredResources. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/preview/2021-09-01-preview/examples/Monitors_ListMonitoredResources_MinimumSet_Gen.json
     */
    /**
     * Sample code: Monitors_ListMonitoredResources_MinimumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void monitorsListMonitoredResourcesMinimumSetGen(
        com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager.monitors().listMonitoredResources("myResourceGroup", "myMonitor", Context.NONE);
    }
}
