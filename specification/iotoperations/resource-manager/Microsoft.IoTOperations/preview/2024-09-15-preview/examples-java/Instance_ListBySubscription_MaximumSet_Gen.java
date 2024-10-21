
/**
 * Samples for Instance List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-15-preview/Instance_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: Instance_ListBySubscription.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void
        instanceListBySubscription(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.instances().list(com.azure.core.util.Context.NONE);
    }
}
