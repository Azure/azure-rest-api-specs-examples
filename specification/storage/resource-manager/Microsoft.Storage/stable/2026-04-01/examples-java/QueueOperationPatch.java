
import com.azure.resourcemanager.storage.fluent.models.StorageQueueInner;

/**
 * Samples for Queue Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/QueueOperationPatch.json
     */
    /**
     * Sample code: QueueOperationPatch.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void queueOperationPatch(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getQueues().updateWithResponse("res3376", "sto328", "queue6185",
            new StorageQueueInner(), com.azure.core.util.Context.NONE);
    }
}
