/** Samples for Firmware Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/preview/2023-02-08-preview/examples/Firmware_Create_MinimumSet_Gen.json
     */
    /**
     * Sample code: Firmware_Create_MinimumSet_Gen.
     *
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void firmwareCreateMinimumSetGen(
        com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        manager.firmwares().define("umrkdttp").withExistingWorkspace("rgworkspaces-firmwares", "A7").create();
    }
}
