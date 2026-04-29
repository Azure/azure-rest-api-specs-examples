
/**
 * Samples for VirtualMachineImages ListOffers.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/virtualMachineImageExamples/VirtualMachineImage_ListOffers_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineImage_ListOffers_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineImageListOffersMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineImages().listOffersWithResponse("aaaaaaa", "aaaaaaaa",
            com.azure.core.util.Context.NONE);
    }
}
