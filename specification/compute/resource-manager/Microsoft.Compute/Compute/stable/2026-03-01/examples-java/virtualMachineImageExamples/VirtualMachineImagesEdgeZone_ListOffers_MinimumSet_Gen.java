
/**
 * Samples for VirtualMachineImagesEdgeZone ListOffers.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-03-01/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_ListOffers_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineImagesEdgeZone_ListOffers_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineImagesEdgeZoneListOffersMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineImagesEdgeZones().listOffersWithResponse("aaaaaaaaaaaaaa",
            "aaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
