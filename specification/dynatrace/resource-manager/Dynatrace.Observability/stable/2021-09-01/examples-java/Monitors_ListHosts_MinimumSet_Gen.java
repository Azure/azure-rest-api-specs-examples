import com.azure.core.util.Context;

/** Samples for Monitors ListHosts. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2021-09-01/examples/Monitors_ListHosts_MinimumSet_Gen.json
     */
    /**
     * Sample code: Monitors_ListHosts_MinimumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void monitorsListHostsMinimumSetGen(com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager.monitors().listHosts("myResourceGroup", "myMonitor", Context.NONE);
    }
}
