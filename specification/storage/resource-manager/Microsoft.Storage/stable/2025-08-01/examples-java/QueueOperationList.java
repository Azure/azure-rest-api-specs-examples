
/**
 * Samples for Queue List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/QueueOperationList.json
     */
    /**
     * Sample code: QueueOperationList.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void queueOperationList(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getQueues().list("res9290", "sto328", null, null, com.azure.core.util.Context.NONE);
    }
}
