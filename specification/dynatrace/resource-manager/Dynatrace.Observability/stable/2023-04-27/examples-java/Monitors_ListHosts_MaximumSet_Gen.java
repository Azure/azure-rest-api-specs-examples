/** Samples for Monitors ListHosts. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2023-04-27/examples/Monitors_ListHosts_MaximumSet_Gen.json
     */
    /**
     * Sample code: Monitors_ListHosts_MaximumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void monitorsListHostsMaximumSetGen(com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager.monitors().listHosts("myResourceGroup", "myMonitor", com.azure.core.util.Context.NONE);
    }
}
