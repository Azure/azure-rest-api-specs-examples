
/**
 * Samples for VirtualMachineImages ListOffers.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/
     * virtualMachineImageExamples/VirtualMachineImage_ListOffers_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineImage_ListOffers_MaximumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        virtualMachineImageListOffersMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineImages().listOffersWithResponse("aaaaaaa",
            "aaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
