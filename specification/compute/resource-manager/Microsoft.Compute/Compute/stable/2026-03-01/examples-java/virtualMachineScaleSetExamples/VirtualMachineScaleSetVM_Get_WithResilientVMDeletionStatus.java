
/**
 * Samples for VirtualMachineScaleSetVMs Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_Get_WithResilientVMDeletionStatus.json
     */
    /**
     * Sample code: Get VM scale set VM with ResiliencyView.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getVMScaleSetVMWithResiliencyView(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetVMs().getWithResponse("myResourceGroup", "{vmss-name}", "1",
            null, com.azure.core.util.Context.NONE);
    }
}
