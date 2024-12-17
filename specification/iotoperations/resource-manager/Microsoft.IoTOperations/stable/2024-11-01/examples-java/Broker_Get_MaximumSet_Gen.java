
/**
 * Samples for Broker Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/Broker_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Broker_Get.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void brokerGet(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.brokers().getWithResponse("rgiotoperations", "resource-name123", "resource-name123",
            com.azure.core.util.Context.NONE);
    }
}
