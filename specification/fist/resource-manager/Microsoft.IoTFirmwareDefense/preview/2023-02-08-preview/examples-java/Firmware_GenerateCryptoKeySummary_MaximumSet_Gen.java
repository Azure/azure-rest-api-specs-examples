/** Samples for Firmware GenerateCryptoKeySummary. */
public final class Main {
    /*
     * x-ms-original-file: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/preview/2023-02-08-preview/examples/Firmware_GenerateCryptoKeySummary_MaximumSet_Gen.json
     */
    /**
     * Sample code: Firmware_GenerateCryptoKeySummary_MaximumSet_Gen.
     *
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void firmwareGenerateCryptoKeySummaryMaximumSetGen(
        com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        manager
            .firmwares()
            .generateCryptoKeySummaryWithResponse(
                "FirmwareAnalysisRG",
                "default",
                "DECAFBAD-0000-0000-0000-BADBADBADBAD",
                com.azure.core.util.Context.NONE);
    }
}
