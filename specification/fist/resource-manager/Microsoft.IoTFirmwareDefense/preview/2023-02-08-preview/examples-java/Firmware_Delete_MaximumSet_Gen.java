/** Samples for Firmware Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/preview/2023-02-08-preview/examples/Firmware_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Firmware_Delete_MaximumSet_Gen.
     *
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void firmwareDeleteMaximumSetGen(
        com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        manager
            .firmwares()
            .deleteWithResponse("rgworkspaces-firmwares", "A7", "umrkdttp", com.azure.core.util.Context.NONE);
    }
}
