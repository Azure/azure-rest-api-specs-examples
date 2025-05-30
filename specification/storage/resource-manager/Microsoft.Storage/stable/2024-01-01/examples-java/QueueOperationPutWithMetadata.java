
import com.azure.resourcemanager.storage.fluent.models.StorageQueueInner;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Queue Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/QueueOperationPutWithMetadata
     * .json
     */
    /**
     * Sample code: QueueOperationPutWithMetadata.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void queueOperationPutWithMetadata(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getQueues().createWithResponse("res3376", "sto328",
            "queue6185", new StorageQueueInner().withMetadata(mapOf("sample1", "meta1", "sample2", "meta2")),
            com.azure.core.util.Context.NONE);
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
