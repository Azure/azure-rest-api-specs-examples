/** Samples for Monitors ListAppServices. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2023-04-27/examples/Monitors_ListAppServices_MaximumSet_Gen.json
     */
    /**
     * Sample code: Monitors_ListAppServices_MaximumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void monitorsListAppServicesMaximumSetGen(
        com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager.monitors().listAppServices("myResourceGroup", "myMonitor", com.azure.core.util.Context.NONE);
    }
}
