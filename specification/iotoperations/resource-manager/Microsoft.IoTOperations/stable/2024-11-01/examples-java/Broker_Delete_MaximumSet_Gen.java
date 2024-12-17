
/**
 * Samples for Broker Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/Broker_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Broker_Delete.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void brokerDelete(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.brokers().delete("rgiotoperations", "resource-name123", "resource-name123",
            com.azure.core.util.Context.NONE);
    }
}
