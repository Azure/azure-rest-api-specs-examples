
/**
 * Samples for VirtualMachineImagesEdgeZone Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_Get_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineImagesEdgeZone_Get_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineImagesEdgeZoneGetMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineImagesEdgeZones().getWithResponse("aaaaaaaaaaaaaaaaaaaaaaa",
            "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaa", "aaaaaaaaaaaaaaaaaa", "aa",
            com.azure.core.util.Context.NONE);
    }
}
