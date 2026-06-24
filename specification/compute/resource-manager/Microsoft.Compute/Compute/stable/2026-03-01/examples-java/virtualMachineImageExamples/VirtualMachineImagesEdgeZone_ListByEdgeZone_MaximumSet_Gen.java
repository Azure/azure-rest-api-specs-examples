
/**
 * Samples for VirtualMachineImages ListByEdgeZone.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-03-01/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_ListByEdgeZone_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineImagesEdgeZone_ListByEdgeZone_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineImagesEdgeZoneListByEdgeZoneMaximumSetGen(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineImages().listByEdgeZoneWithResponse("WestUS", "microsoftlosangeles1",
            com.azure.core.util.Context.NONE);
    }
}
