import com.azure.core.util.Context;

/** Samples for SapVirtualInstances ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/sapvirtualinstances/SAPVirtualInstances_ListByResourceGroup.json
     */
    /**
     * Sample code: SAPVirtualInstances_ListByResourceGroup.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void sAPVirtualInstancesListByResourceGroup(
        com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager.sapVirtualInstances().listByResourceGroup("test-rg", Context.NONE);
    }
}
