
/**
 * Samples for Machines GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-15/machine/Machines_Get.json
     */
    /**
     * Sample code: Get Machine.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void getMachine(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.machines().getByResourceGroupWithResponse("myResourceGroup", "myMachine", null,
            com.azure.core.util.Context.NONE);
    }
}
