/** Samples for VirtualMachineImages ListSkus. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/examples/virtualMachineImageExamples/VirtualMachineImage_ListSkus_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineImage_ListSkus_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineImageListSkusMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineImages()
            .listSkusWithResponse("aaaa", "aaaaaaaaaaaaa", "aaaaaaa", com.azure.core.util.Context.NONE);
    }
}
