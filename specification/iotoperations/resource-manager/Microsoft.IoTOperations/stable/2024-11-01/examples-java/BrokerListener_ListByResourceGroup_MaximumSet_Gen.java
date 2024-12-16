
/**
 * Samples for BrokerListener ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/BrokerListener_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: BrokerListener_ListByResourceGroup.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void
        brokerListenerListByResourceGroup(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.brokerListeners().listByResourceGroup("rgiotoperations", "resource-name123", "resource-name123",
            com.azure.core.util.Context.NONE);
    }
}
