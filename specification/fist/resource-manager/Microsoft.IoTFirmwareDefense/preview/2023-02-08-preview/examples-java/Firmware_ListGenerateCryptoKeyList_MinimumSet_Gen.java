/** Samples for Firmware ListGenerateCryptoKeyList. */
public final class Main {
    /*
     * x-ms-original-file: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/preview/2023-02-08-preview/examples/Firmware_ListGenerateCryptoKeyList_MinimumSet_Gen.json
     */
    /**
     * Sample code: Firmware_ListGenerateCryptoKeyList_MinimumSet_Gen.
     *
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void firmwareListGenerateCryptoKeyListMinimumSetGen(
        com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        manager
            .firmwares()
            .listGenerateCryptoKeyList(
                "rgworkspaces-firmwares", "j5QE_", "wujtpcgypfpqseyrsebolarkspy", com.azure.core.util.Context.NONE);
    }
}
