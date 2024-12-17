
/**
 * Samples for BrokerListener Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/BrokerListener_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: BrokerListener_Delete.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void brokerListenerDelete(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.brokerListeners().delete("rgiotoperations", "resource-name123", "resource-name123", "resource-name123",
            com.azure.core.util.Context.NONE);
    }
}
