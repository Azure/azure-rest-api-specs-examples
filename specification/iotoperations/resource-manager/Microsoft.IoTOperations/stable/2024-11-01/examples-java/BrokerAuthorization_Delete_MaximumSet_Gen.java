
/**
 * Samples for BrokerAuthorization Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/BrokerAuthorization_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: BrokerAuthorization_Delete.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void brokerAuthorizationDelete(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.brokerAuthorizations().delete("rgiotoperations", "resource-name123", "resource-name123",
            "resource-name123", com.azure.core.util.Context.NONE);
    }
}
