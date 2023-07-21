/** Samples for Firmware GenerateSummary. */
public final class Main {
    /*
     * x-ms-original-file: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/preview/2023-02-08-preview/examples/Firmware_GenerateSummary_MaximumSet_Gen.json
     */
    /**
     * Sample code: Firmware_GenerateSummary_MaximumSet_Gen.
     *
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void firmwareGenerateSummaryMaximumSetGen(
        com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        manager
            .firmwares()
            .generateSummaryWithResponse("rgworkspaces-firmwares", "A7", "umrkdttp", com.azure.core.util.Context.NONE);
    }
}
