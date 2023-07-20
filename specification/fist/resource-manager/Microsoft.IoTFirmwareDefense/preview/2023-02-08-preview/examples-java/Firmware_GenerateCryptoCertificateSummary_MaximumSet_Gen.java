/** Samples for Firmware GenerateCryptoCertificateSummary. */
public final class Main {
    /*
     * x-ms-original-file: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/preview/2023-02-08-preview/examples/Firmware_GenerateCryptoCertificateSummary_MaximumSet_Gen.json
     */
    /**
     * Sample code: Firmware_GenerateCryptoCertificateSummary_MaximumSet_Gen.
     *
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void firmwareGenerateCryptoCertificateSummaryMaximumSetGen(
        com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        manager
            .firmwares()
            .generateCryptoCertificateSummaryWithResponse(
                "FirmwareAnalysisRG",
                "default",
                "DECAFBAD-0000-0000-0000-BADBADBADBAD",
                com.azure.core.util.Context.NONE);
    }
}
