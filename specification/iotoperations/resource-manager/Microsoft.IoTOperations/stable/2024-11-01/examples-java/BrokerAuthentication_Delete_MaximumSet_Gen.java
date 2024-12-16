
/**
 * Samples for BrokerAuthentication Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/BrokerAuthentication_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: BrokerAuthentication_Delete.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void
        brokerAuthenticationDelete(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.brokerAuthentications().delete("rgiotoperations", "resource-name123", "resource-name123",
            "resource-name123", com.azure.core.util.Context.NONE);
    }
}
