
/**
 * Samples for VirtualMachineImages Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/virtualMachineImageExamples/VirtualMachineImage_Get_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineImage_Get_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineImageGetMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineImages().getWithResponse("aaaaaaaaaaaa", "aaaaaaaaaaa", "aa",
            "aaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
