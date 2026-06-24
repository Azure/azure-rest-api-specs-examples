
/**
 * Samples for VirtualMachineImages List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineImageExamples/VirtualMachineImage_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineImage_List_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineImageListMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineImages().listWithResponse("aaaaaaaaaaaaaaa", "aaaaaa",
            "aaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaa", 18, "aa",
            com.azure.core.util.Context.NONE);
    }
}
