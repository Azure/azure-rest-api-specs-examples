/** Samples for Firmware GenerateBinaryHardeningSummary. */
public final class Main {
    /*
     * x-ms-original-file: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/preview/2023-02-08-preview/examples/Firmware_GenerateBinaryHardeningSummary_MinimumSet_Gen.json
     */
    /**
     * Sample code: Firmware_GenerateBinaryHardeningSummary_MinimumSet_Gen.
     *
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void firmwareGenerateBinaryHardeningSummaryMinimumSetGen(
        com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        manager
            .firmwares()
            .generateBinaryHardeningSummaryWithResponse(
                "rgworkspaces-firmwares", "A7", "umrkdttp", com.azure.core.util.Context.NONE);
    }
}
