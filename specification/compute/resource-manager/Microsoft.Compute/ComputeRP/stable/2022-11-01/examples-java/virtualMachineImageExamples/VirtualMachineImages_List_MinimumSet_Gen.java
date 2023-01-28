/** Samples for VirtualMachineImages List. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-11-01/examples/virtualMachineImageExamples/VirtualMachineImages_List_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineImages_List_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineImagesListMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineImages()
            .listWithResponse(
                "aaaaaaa", "aaaaaaaaaaa", "aaaaaaaaaa", "aaaaaa", null, null, null, com.azure.core.util.Context.NONE);
    }
}
