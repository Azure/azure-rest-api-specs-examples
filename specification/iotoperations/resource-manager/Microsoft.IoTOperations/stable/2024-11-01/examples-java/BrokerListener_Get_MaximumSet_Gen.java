
/**
 * Samples for BrokerListener Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/BrokerListener_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: BrokerListener_Get.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void brokerListenerGet(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.brokerListeners().getWithResponse("rgiotoperations", "resource-name123", "resource-name123",
            "resource-name123", com.azure.core.util.Context.NONE);
    }
}
