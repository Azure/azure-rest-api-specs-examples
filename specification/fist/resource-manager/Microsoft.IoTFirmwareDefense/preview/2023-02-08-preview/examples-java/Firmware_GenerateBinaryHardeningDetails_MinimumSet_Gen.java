/** Samples for Firmware GenerateBinaryHardeningDetails. */
public final class Main {
    /*
     * x-ms-original-file: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/preview/2023-02-08-preview/examples/Firmware_GenerateBinaryHardeningDetails_MinimumSet_Gen.json
     */
    /**
     * Sample code: Firmware_GenerateBinaryHardeningDetails_MinimumSet_Gen.
     *
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void firmwareGenerateBinaryHardeningDetailsMinimumSetGen(
        com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        manager
            .firmwares()
            .generateBinaryHardeningDetailsWithResponse(
                "rgworkspaces-firmwares", "A7", "umrkdttp", com.azure.core.util.Context.NONE);
    }
}
