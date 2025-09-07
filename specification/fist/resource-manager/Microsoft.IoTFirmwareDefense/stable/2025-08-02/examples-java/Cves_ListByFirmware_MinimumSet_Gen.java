
/**
 * Samples for Cves ListByFirmware.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-02/Cves_ListByFirmware_MinimumSet_Gen.json
     */
    /**
     * Sample code: Cves_ListByFirmware_MinimumSet_Gen.
     * 
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void cvesListByFirmwareMinimumSetGen(
        com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        manager.cves().listByFirmware("FirmwareAnalysisRG", "default", "00000000-0000-0000-0000-000000000000",
            com.azure.core.util.Context.NONE);
    }
}
