
/**
 * Samples for VirtualMachineExtensionImages Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2025-11-01/virtualMachineExtensionImageExamples/VirtualMachineExtensionImage_Get_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineExtensionImage_Get_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineExtensionImageGetMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineExtensionImages().getWithResponse("aaaaaaaaaaaaaa",
            "aaaaaaaaaaaaaaaaaaaaaaaaaa", "aa", "aaa", com.azure.core.util.Context.NONE);
    }
}
