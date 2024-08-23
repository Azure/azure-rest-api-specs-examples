
/**
 * Samples for VirtualMachineImages List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/
     * virtualMachineImageExamples/VirtualMachineImage_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineImage_List_MaximumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineImageListMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineImages().listWithResponse("aaaaaaaaaaaaaaa",
            "aaaaaa", "aaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaa", 18, "aa",
            com.azure.core.util.Context.NONE);
    }
}
