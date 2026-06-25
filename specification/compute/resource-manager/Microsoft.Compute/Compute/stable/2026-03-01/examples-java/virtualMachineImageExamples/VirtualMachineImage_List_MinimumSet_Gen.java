
/**
 * Samples for VirtualMachineImages List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineImageExamples/VirtualMachineImage_List_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineImage_List_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineImageListMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineImages().listWithResponse("aaaaaaa", "aaaaaaaaaaa", "aaaaaaaaaa",
            "aaaaaa", null, null, null, com.azure.core.util.Context.NONE);
    }
}
