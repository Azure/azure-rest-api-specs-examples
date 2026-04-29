
/**
 * Samples for VirtualMachineImagesEdgeZone ListPublishers.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2025-11-01/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_ListPublishers_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineImagesEdgeZone_ListPublishers_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineImagesEdgeZoneListPublishersMinimumSetGen(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineImagesEdgeZones().listPublishersWithResponse("aaaa",
            "aaaaaaaaaaaaaaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
