
/**
 * Samples for VirtualMachineExtensionImages Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2025-11-01/virtualMachineExtensionImageExamples/VirtualMachineExtensionImage_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineExtensionImage_Get_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineExtensionImageGetMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineExtensionImages().getWithResponse("aaaaaaaaaaaaa",
            "aaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
