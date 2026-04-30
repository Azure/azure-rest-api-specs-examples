
/**
 * Samples for VirtualMachineScaleSetVMs Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_Get_WithUserData.json
     */
    /**
     * Sample code: Get VM scale set VM with UserData.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getVMScaleSetVMWithUserData(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetVMs().getWithResponse("myResourceGroup", "{vmss-name}", "0",
            null, com.azure.core.util.Context.NONE);
    }
}
