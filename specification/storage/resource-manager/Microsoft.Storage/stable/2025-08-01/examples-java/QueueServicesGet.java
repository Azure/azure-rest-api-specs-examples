
/**
 * Samples for QueueServices GetServiceProperties.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/QueueServicesGet.json
     */
    /**
     * Sample code: QueueServicesGet.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void queueServicesGet(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getQueueServices().getServicePropertiesWithResponse("res4410", "sto8607",
            com.azure.core.util.Context.NONE);
    }
}
