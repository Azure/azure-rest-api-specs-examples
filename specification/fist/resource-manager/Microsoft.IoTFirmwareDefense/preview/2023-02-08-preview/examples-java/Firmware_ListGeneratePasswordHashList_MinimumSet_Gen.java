/** Samples for Firmware ListGeneratePasswordHashList. */
public final class Main {
    /*
     * x-ms-original-file: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/preview/2023-02-08-preview/examples/Firmware_ListGeneratePasswordHashList_MinimumSet_Gen.json
     */
    /**
     * Sample code: Firmware_ListGeneratePasswordHashList_MinimumSet_Gen.
     *
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void firmwareListGeneratePasswordHashListMinimumSetGen(
        com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        manager
            .firmwares()
            .listGeneratePasswordHashList("rgworkspaces-firmwares", "A7", "umrkdttp", com.azure.core.util.Context.NONE);
    }
}
