
/**
 * Samples for AkriConnector Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/AkriConnector_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: AkriConnector_Get_MaximumSet.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void
        akriConnectorGetMaximumSet(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.akriConnectors().getWithResponse("rgiotoperations", "resource-name123", "resource-name123",
            "resource-name123", com.azure.core.util.Context.NONE);
    }
}
