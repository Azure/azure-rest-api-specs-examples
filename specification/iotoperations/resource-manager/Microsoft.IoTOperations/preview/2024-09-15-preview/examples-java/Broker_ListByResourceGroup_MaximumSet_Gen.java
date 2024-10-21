
/**
 * Samples for Broker ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-15-preview/Broker_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: Broker_ListByResourceGroup.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void brokerListByResourceGroup(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.brokers().listByResourceGroup("rgiotoperations", "resource-name123", com.azure.core.util.Context.NONE);
    }
}
