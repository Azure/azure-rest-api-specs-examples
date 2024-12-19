
/**
 * Samples for Firmwares Create.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/stable/2024-01-10/examples/
     * Firmwares_Create_MinimumSet_Gen.json
     */
    /**
     * Sample code: Firmwares_Create_MinimumSet_Gen.
     * 
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void
        firmwaresCreateMinimumSetGen(com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        manager.firmwares().define("umrkdttp").withExistingWorkspace("rgworkspaces-firmwares", "A7").create();
    }
}
