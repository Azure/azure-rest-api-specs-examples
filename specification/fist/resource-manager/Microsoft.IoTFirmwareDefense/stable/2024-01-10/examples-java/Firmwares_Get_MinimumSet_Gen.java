
/**
 * Samples for Firmwares Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/stable/2024-01-10/examples/
     * Firmwares_Get_MinimumSet_Gen.json
     */
    /**
     * Sample code: Firmwares_Get_MinimumSet_Gen.
     * 
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void
        firmwaresGetMinimumSetGen(com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        manager.firmwares().getWithResponse("rgworkspaces-firmwares", "A7", "umrkdttp",
            com.azure.core.util.Context.NONE);
    }
}
