/** Samples for Monitors Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2023-04-27/examples/Monitors_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: Monitors_Delete_MinimumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void monitorsDeleteMinimumSetGen(com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager.monitors().delete("myResourceGroup", "myMonitor", com.azure.core.util.Context.NONE);
    }
}
