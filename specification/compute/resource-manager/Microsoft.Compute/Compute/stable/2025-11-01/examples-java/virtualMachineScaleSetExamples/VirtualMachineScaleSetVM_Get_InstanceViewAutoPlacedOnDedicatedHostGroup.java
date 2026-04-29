
/**
 * Samples for VirtualMachineScaleSetVMs GetInstanceView.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_Get_InstanceViewAutoPlacedOnDedicatedHostGroup
     * .json
     */
    /**
     * Sample code: Get instance view of a virtual machine from a VM scale set placed on a dedicated host group through
     * automatic placement.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        getInstanceViewOfAVirtualMachineFromAVMScaleSetPlacedOnADedicatedHostGroupThroughAutomaticPlacement(
            com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetVMs().getInstanceViewWithResponse("myResourceGroup",
            "myVirtualMachineScaleSet", "0", com.azure.core.util.Context.NONE);
    }
}
