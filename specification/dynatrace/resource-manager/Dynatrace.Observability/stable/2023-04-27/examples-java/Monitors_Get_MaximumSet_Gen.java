/** Samples for Monitors GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2023-04-27/examples/Monitors_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Monitors_Get_MaximumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void monitorsGetMaximumSetGen(com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager
            .monitors()
            .getByResourceGroupWithResponse("myResourceGroup", "myMonitor", com.azure.core.util.Context.NONE);
    }
}
