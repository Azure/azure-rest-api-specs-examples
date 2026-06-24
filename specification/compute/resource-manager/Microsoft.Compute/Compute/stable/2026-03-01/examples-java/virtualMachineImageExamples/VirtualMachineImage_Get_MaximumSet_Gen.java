
/**
 * Samples for VirtualMachineImages Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineImageExamples/VirtualMachineImage_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineImage_Get_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineImageGetMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineImages().getWithResponse("aaaaaa", "aaa",
            "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaa",
            com.azure.core.util.Context.NONE);
    }
}
