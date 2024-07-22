
/**
 * Samples for Machines ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-05-20-preview/examples/machine/
     * Machines_ListByResourceGroup.json
     */
    /**
     * Sample code: List Machines by resource group.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void
        listMachinesByResourceGroup(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.machines().listByResourceGroup("myResourceGroup", null, com.azure.core.util.Context.NONE);
    }
}
