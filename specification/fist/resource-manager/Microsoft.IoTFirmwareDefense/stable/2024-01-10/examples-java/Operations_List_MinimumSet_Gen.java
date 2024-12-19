
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/stable/2024-01-10/examples/
     * Operations_List_MinimumSet_Gen.json
     */
    /**
     * Sample code: Operations_List_MinimumSet_Gen.
     * 
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void
        operationsListMinimumSetGen(com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
