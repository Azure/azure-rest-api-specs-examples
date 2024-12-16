
/**
 * Samples for BrokerAuthorization ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/BrokerAuthorization_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: BrokerAuthorization_ListByResourceGroup.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void
        brokerAuthorizationListByResourceGroup(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.brokerAuthorizations().listByResourceGroup("rgiotoperations", "resource-name123", "resource-name123",
            com.azure.core.util.Context.NONE);
    }
}
