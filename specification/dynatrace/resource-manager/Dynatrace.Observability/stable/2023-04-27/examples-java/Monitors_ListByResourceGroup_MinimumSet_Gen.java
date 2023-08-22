/** Samples for Monitors ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2023-04-27/examples/Monitors_ListByResourceGroup_MinimumSet_Gen.json
     */
    /**
     * Sample code: Monitors_ListByResourceGroup_MinimumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void monitorsListByResourceGroupMinimumSetGen(
        com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager.monitors().listByResourceGroup("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
