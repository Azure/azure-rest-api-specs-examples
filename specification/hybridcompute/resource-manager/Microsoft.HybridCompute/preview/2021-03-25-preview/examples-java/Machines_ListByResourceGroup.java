/** Samples for Machines ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2021-03-25-preview/examples/Machines_ListByResourceGroup.json
     */
    /**
     * Sample code: List Machines by resource group.
     *
     * @param manager Entry point to HybridComputeManager.
     */
    public static void listMachinesByResourceGroup(
        com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.machines().listByResourceGroup("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
