import com.azure.core.util.Context;

/** Samples for VirtualMachineImages ListSkus. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineImageExamples/VirtualMachineImages_ListSkus_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineImages_ListSkus_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineImagesListSkusMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineImages()
            .listSkusWithResponse("aaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaa", Context.NONE);
    }
}
