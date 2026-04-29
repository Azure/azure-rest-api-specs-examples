
/**
 * Samples for VirtualMachines GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/virtualMachineExamples/VirtualMachine_Get_AutoPlacedOnDedicatedHostGroup.json
     */
    /**
     * Sample code: Get a virtual machine placed on a dedicated host group through automatic placement.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getAVirtualMachinePlacedOnADedicatedHostGroupThroughAutomaticPlacement(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().getByResourceGroupWithResponse("myResourceGroup", "myVM", null,
            com.azure.core.util.Context.NONE);
    }
}
