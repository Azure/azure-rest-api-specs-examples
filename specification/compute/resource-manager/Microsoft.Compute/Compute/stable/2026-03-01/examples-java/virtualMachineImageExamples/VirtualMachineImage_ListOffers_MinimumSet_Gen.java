
/**
 * Samples for VirtualMachineImages ListOffers.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineImageExamples/VirtualMachineImage_ListOffers_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineImage_ListOffers_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineImageListOffersMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineImages().listOffersWithResponse("aaaaaaaaaaaaaaaaa",
            "aaaaaaaaaaaaaaaaaaaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
