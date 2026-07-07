
/**
 * Samples for Queue Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/QueueOperationDelete.json
     */
    /**
     * Sample code: QueueOperationDelete.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void queueOperationDelete(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getQueues().deleteWithResponse("res3376", "sto328", "queue6185",
            com.azure.core.util.Context.NONE);
    }
}
