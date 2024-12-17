
/**
 * Samples for BrokerAuthentication ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/BrokerAuthentication_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: BrokerAuthentication_ListByResourceGroup.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void
        brokerAuthenticationListByResourceGroup(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.brokerAuthentications().listByResourceGroup("rgiotoperations", "resource-name123", "resource-name123",
            com.azure.core.util.Context.NONE);
    }
}
