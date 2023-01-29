/** Samples for VirtualMachineImages ListSkus. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-11-01/examples/virtualMachineImageExamples/VirtualMachineImages_ListSkus_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineImages_ListSkus_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineImagesListSkusMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineImages()
            .listSkusWithResponse("aaaa", "aaaaaaaaaaaaa", "aaaaaaa", com.azure.core.util.Context.NONE);
    }
}
