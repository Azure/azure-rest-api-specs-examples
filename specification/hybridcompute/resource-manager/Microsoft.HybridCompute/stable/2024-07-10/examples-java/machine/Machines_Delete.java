
/**
 * Samples for Machines Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridcompute/resource-manager/Microsoft.HybridCompute/stable/2024-07-10/examples/machine/
     * Machines_Delete.json
     */
    /**
     * Sample code: Delete a Machine.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void deleteAMachine(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.machines().deleteByResourceGroupWithResponse("myResourceGroup", "myMachine",
            com.azure.core.util.Context.NONE);
    }
}
