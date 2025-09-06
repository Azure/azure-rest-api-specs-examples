
/**
 * Samples for BinaryHardening ListByFirmware.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-02/BinaryHardening_ListByFirmware_MinimumSet_Gen.json
     */
    /**
     * Sample code: BinaryHardening_ListByFirmware_MinimumSet_Gen.
     * 
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void binaryHardeningListByFirmwareMinimumSetGen(
        com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        manager.binaryHardenings().listByFirmware("FirmwareAnalysisRG", "default",
            "00000000-0000-0000-0000-000000000000", com.azure.core.util.Context.NONE);
    }
}
