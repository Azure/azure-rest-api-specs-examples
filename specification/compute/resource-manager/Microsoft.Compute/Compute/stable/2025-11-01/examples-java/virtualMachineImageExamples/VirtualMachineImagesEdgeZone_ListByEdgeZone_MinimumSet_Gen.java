
/**
 * Samples for VirtualMachineImages ListByEdgeZone.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2025-11-01/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_ListByEdgeZone_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineImagesEdgeZone_ListByEdgeZone_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineImagesEdgeZoneListByEdgeZoneMinimumSetGen(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineImages().listByEdgeZoneWithResponse("WestUS", "microsoftlosangeles1",
            com.azure.core.util.Context.NONE);
    }
}
