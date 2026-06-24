
/**
 * Samples for VirtualMachineImagesEdgeZone ListPublishers.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-03-01/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_ListPublishers_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineImagesEdgeZone_ListPublishers_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineImagesEdgeZoneListPublishersMaximumSetGen(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineImagesEdgeZones().listPublishersWithResponse("aaaaaa",
            "aaaaaaaaaaaaaaaaaaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
