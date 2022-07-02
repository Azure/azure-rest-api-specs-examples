import com.azure.core.util.Context;

/** Samples for SapVirtualInstances GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/sapvirtualinstances/SAPVirtualInstances_Get.json
     */
    /**
     * Sample code: SAPVirtualInstances_Get.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void sAPVirtualInstancesGet(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager.sapVirtualInstances().getByResourceGroupWithResponse("test-rg", "X00", Context.NONE);
    }
}
