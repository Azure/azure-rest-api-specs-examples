
import com.azure.resourcemanager.storage.fluent.models.StorageQueueInner;

/**
 * Samples for Queue Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/QueueOperationPatch.json
     */
    /**
     * Sample code: QueueOperationPatch.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void queueOperationPatch(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getQueues().updateWithResponse("res3376", "sto328",
            "queue6185", new StorageQueueInner(), com.azure.core.util.Context.NONE);
    }
}
