
/**
 * Samples for BinaryHardening ListByFirmware.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/stable/2024-01-10/examples/
     * BinaryHardening_ListByFirmware_MaximumSet_Gen.json
     */
    /**
     * Sample code: BinaryHardening_ListByFirmware_MaximumSet_Gen.
     * 
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void binaryHardeningListByFirmwareMaximumSetGen(
        com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        manager.binaryHardenings().listByFirmware("FirmwareAnalysisRG", "default",
            "109a9886-50bf-85a8-9d75-000000000000", com.azure.core.util.Context.NONE);
    }
}
