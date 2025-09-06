
/**
 * Samples for PasswordHashes ListByFirmware.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-02/PasswordHashes_ListByFirmware_MinimumSet_Gen.json
     */
    /**
     * Sample code: PasswordHashes_ListByFirmware_MinimumSet_Gen.
     * 
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void passwordHashesListByFirmwareMinimumSetGen(
        com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        manager.passwordHashes().listByFirmware("FirmwareAnalysisRG", "default", "00000000-0000-0000-0000-000000000000",
            com.azure.core.util.Context.NONE);
    }
}
