
/**
 * Samples for VirtualMachineScaleSetVMExtensions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVMExtension_Get.json
     */
    /**
     * Sample code: Get VirtualMachineScaleSet VM extension.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getVirtualMachineScaleSetVMExtension(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetVMExtensions().getWithResponse("myResourceGroup",
            "myvmScaleSet", "0", "myVMExtension", null, com.azure.core.util.Context.NONE);
    }
}
