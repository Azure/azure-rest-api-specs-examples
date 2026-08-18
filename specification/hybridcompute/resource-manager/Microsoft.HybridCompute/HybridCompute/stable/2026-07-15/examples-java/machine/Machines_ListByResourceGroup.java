
/**
 * Samples for Machines ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-15/machine/Machines_ListByResourceGroup.json
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
