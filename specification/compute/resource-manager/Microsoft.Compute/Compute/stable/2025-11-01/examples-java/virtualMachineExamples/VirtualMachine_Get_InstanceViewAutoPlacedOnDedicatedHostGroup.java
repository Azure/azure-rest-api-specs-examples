
/**
 * Samples for VirtualMachines InstanceView.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2025-11-01/virtualMachineExamples/VirtualMachine_Get_InstanceViewAutoPlacedOnDedicatedHostGroup.json
     */
    /**
     * Sample code: Get instance view of a virtual machine placed on a dedicated host group through automatic placement.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getInstanceViewOfAVirtualMachinePlacedOnADedicatedHostGroupThroughAutomaticPlacement(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().instanceViewWithResponse("myResourceGroup", "myVM",
            com.azure.core.util.Context.NONE);
    }
}
