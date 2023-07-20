/** Samples for Firmware GenerateCryptoKeySummary. */
public final class Main {
    /*
     * x-ms-original-file: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/preview/2023-02-08-preview/examples/Firmware_GenerateCryptoKeySummary_MinimumSet_Gen.json
     */
    /**
     * Sample code: Firmware_GenerateCryptoKeySummary_MinimumSet_Gen.
     *
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void firmwareGenerateCryptoKeySummaryMinimumSetGen(
        com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        manager
            .firmwares()
            .generateCryptoKeySummaryWithResponse(
                "rgworkspaces-firmwares", "j5QE_", "wujtpcgypfpqseyrsebolarkspy", com.azure.core.util.Context.NONE);
    }
}
