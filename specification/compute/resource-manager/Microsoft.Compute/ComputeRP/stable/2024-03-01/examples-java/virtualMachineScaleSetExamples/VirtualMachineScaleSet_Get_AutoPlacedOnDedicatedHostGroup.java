
/**
 * Samples for VirtualMachineScaleSets GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSet_Get_AutoPlacedOnDedicatedHostGroup.json
     */
    /**
     * Sample code: Get a virtual machine scale set placed on a dedicated host group through automatic placement.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAVirtualMachineScaleSetPlacedOnADedicatedHostGroupThroughAutomaticPlacement(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSets().getByResourceGroupWithResponse(
            "myResourceGroup", "myVirtualMachineScaleSet", null, com.azure.core.util.Context.NONE);
    }
}
