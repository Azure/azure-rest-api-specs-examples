
/**
 * Samples for VirtualMachineScaleSets GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Get.json
     */
    /**
     * Sample code: Get a virtual machine scale set.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getAVirtualMachineScaleSet(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSets().getByResourceGroupWithResponse("myResourceGroup",
            "myVirtualMachineScaleSet", null, com.azure.core.util.Context.NONE);
    }
}
