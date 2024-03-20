
/**
 * Samples for Firmwares ListByWorkspace.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/stable/2024-01-10/examples/
     * Firmwares_ListByWorkspace_MaximumSet_Gen.json
     */
    /**
     * Sample code: Firmwares_ListByWorkspace_MaximumSet_Gen.
     * 
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void firmwaresListByWorkspaceMaximumSetGen(
        com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        manager.firmwares().listByWorkspace("rgworkspaces-firmwares", "A7", com.azure.core.util.Context.NONE);
    }
}
