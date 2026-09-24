
/**
 * Samples for Queue Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/QueueOperationGet.json
     */
    /**
     * Sample code: QueueOperationGet.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void queueOperationGet(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getQueues().getWithResponse("res3376", "sto328", "queue6185",
            com.azure.core.util.Context.NONE);
    }
}
