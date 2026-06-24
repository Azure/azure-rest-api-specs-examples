
/**
 * Samples for VirtualMachineScaleSets GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Get_AutoPlacedOnDedicatedHostGroup.json
     */
    /**
     * Sample code: Get a virtual machine scale set placed on a dedicated host group through automatic placement.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getAVirtualMachineScaleSetPlacedOnADedicatedHostGroupThroughAutomaticPlacement(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSets().getByResourceGroupWithResponse("myResourceGroup",
            "myVirtualMachineScaleSet", null, com.azure.core.util.Context.NONE);
    }
}
