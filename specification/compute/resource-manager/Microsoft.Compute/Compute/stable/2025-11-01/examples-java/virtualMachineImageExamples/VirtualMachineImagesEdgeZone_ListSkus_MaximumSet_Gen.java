
/**
 * Samples for VirtualMachineImagesEdgeZone ListSkus.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2025-11-01/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_ListSkus_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineImagesEdgeZone_ListSkus_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineImagesEdgeZoneListSkusMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineImagesEdgeZones().listSkusWithResponse("aaaaaaaaaaaa", "aaaaa",
            "aaaaaaaaaaaa", "aaaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
