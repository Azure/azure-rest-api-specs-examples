
/**
 * Samples for CryptoKeys ListByFirmware.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-02/CryptoKeys_ListByFirmware_MinimumSet_Gen.json
     */
    /**
     * Sample code: CryptoKeys_ListByFirmware_MinimumSet_Gen.
     * 
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void cryptoKeysListByFirmwareMinimumSetGen(
        com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        manager.cryptoKeys().listByFirmware("FirmwareAnalysisRG", "default", "00000000-0000-0000-0000-000000000000",
            com.azure.core.util.Context.NONE);
    }
}
