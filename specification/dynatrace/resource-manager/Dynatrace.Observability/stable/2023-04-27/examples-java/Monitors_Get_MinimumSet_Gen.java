/** Samples for Monitors GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2023-04-27/examples/Monitors_Get_MinimumSet_Gen.json
     */
    /**
     * Sample code: Monitors_Get_MinimumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void monitorsGetMinimumSetGen(com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager
            .monitors()
            .getByResourceGroupWithResponse("myResourceGroup", "myMonitor", com.azure.core.util.Context.NONE);
    }
}
