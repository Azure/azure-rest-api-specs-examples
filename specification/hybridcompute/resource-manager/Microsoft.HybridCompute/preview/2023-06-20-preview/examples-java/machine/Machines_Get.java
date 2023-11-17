/** Samples for Machines GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-06-20-preview/examples/machine/Machines_Get.json
     */
    /**
     * Sample code: Get Machine.
     *
     * @param manager Entry point to HybridComputeManager.
     */
    public static void getMachine(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager
            .machines()
            .getByResourceGroupWithResponse("myResourceGroup", "myMachine", null, com.azure.core.util.Context.NONE);
    }
}
