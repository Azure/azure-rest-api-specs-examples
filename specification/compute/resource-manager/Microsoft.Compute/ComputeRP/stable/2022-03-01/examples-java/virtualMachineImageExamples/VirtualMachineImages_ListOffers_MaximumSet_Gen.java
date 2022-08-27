import com.azure.core.util.Context;

/** Samples for VirtualMachineImages ListOffers. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/virtualMachineImageExamples/VirtualMachineImages_ListOffers_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineImages_ListOffers_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineImagesListOffersMaximumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineImages()
            .listOffersWithResponse("aaaaaaa", "aaaaaaaa", Context.NONE);
    }
}
