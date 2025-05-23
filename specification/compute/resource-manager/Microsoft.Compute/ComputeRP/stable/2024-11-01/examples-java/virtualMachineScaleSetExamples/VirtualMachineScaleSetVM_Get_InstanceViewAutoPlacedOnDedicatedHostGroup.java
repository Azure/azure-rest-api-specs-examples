
/**
 * Samples for VirtualMachineScaleSetVMs GetInstanceView.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_Get_InstanceViewAutoPlacedOnDedicatedHostGroup.json
     */
    /**
     * Sample code: Get instance view of a virtual machine from a VM scale set placed on a dedicated host group through
     * automatic placement.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getInstanceViewOfAVirtualMachineFromAVMScaleSetPlacedOnADedicatedHostGroupThroughAutomaticPlacement(
            com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSetVMs().getInstanceViewWithResponse(
            "myResourceGroup", "myVirtualMachineScaleSet", "0", com.azure.core.util.Context.NONE);
    }
}
