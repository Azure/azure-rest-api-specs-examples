
/**
 * Samples for Firmwares ListByWorkspace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-02/Firmwares_ListByWorkspace_MinimumSet_Gen.json
     */
    /**
     * Sample code: Firmwares_ListByWorkspace_MinimumSet_Gen.
     * 
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void firmwaresListByWorkspaceMinimumSetGen(
        com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        manager.firmwares().listByWorkspace("rgworkspaces-firmwares", "A7", com.azure.core.util.Context.NONE);
    }
}
