import com.azure.core.util.Context;

/** Samples for Monitors ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/preview/2021-09-01-preview/examples/Monitors_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: Monitors_ListByResourceGroup_MaximumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void monitorsListByResourceGroupMaximumSetGen(
        com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager.monitors().listByResourceGroup("myResourceGroup", Context.NONE);
    }
}
