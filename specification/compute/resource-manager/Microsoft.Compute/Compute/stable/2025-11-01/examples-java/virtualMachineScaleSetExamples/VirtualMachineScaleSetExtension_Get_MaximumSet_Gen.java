
/**
 * Samples for VirtualMachineScaleSetExtensions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetExtension_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetExtension_Get_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineScaleSetExtensionGetMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetExtensions().getWithResponse("rgcompute",
            "aaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaa", "aaaaaaa", com.azure.core.util.Context.NONE);
    }
}
