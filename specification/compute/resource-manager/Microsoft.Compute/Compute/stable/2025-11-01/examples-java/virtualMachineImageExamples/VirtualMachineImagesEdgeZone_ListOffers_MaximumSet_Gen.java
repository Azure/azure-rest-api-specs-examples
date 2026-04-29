
/**
 * Samples for VirtualMachineImagesEdgeZone ListOffers.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2025-11-01/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_ListOffers_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineImagesEdgeZone_ListOffers_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineImagesEdgeZoneListOffersMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineImagesEdgeZones().listOffersWithResponse("aaaaaaaaaaaaaaaaaa",
            "aaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
