
/**
 * Samples for Machines List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2025-02-19-preview/examples/machine/
     * Machines_ListBySubscription.json
     */
    /**
     * Sample code: List Machines by resource group.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void
        listMachinesByResourceGroup(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.machines().list(com.azure.core.util.Context.NONE);
    }
}
