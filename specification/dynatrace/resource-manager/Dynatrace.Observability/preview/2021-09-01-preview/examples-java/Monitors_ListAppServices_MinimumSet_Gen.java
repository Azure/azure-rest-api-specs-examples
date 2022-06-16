import com.azure.core.util.Context;

/** Samples for Monitors ListAppServices. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/preview/2021-09-01-preview/examples/Monitors_ListAppServices_MinimumSet_Gen.json
     */
    /**
     * Sample code: Monitors_ListAppServices_MinimumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void monitorsListAppServicesMinimumSetGen(
        com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager.monitors().listAppServices("myResourceGroup", "myMonitor", Context.NONE);
    }
}
