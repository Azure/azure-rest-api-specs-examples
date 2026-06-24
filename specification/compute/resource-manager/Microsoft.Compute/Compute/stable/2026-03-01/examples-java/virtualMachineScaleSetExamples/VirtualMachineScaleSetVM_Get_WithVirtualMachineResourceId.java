
/**
 * Samples for VirtualMachineScaleSetVMs Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_Get_WithVirtualMachineResourceId.json
     */
    /**
     * Sample code: Get VM scale set Flex VM with VirtualMachineResourceId.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        getVMScaleSetFlexVMWithVirtualMachineResourceId(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetVMs().getWithResponse("myResourceGroup", "{vmss-flex-name}",
            "{vmss-flex-vm-name}", null, com.azure.core.util.Context.NONE);
    }
}
