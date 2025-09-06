
/**
 * Samples for Firmwares Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-02/Firmwares_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: Firmwares_Delete_MinimumSet_Gen.
     * 
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void
        firmwaresDeleteMinimumSetGen(com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        manager.firmwares().deleteWithResponse("rgworkspaces-firmwares", "A7", "umrkdttp",
            com.azure.core.util.Context.NONE);
    }
}
