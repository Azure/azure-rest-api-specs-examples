
/**
 * Samples for SbomComponents ListByFirmware.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/stable/2024-01-10/examples/
     * SbomComponents_ListByFirmware_MinimumSet_Gen.json
     */
    /**
     * Sample code: SbomComponents_ListByFirmware_MinimumSet_Gen.
     * 
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void sbomComponentsListByFirmwareMinimumSetGen(
        com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        manager.sbomComponents().listByFirmware("FirmwareAnalysisRG", "default", "109a9886-50bf-85a8-9d75-000000000000",
            com.azure.core.util.Context.NONE);
    }
}
