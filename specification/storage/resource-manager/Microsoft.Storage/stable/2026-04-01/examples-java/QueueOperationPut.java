
import com.azure.resourcemanager.storage.fluent.models.StorageQueueInner;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Queue Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/QueueOperationPut.json
     */
    /**
     * Sample code: QueueOperationPut.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void queueOperationPut(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getQueues().createWithResponse("res3376", "sto328", "queue6185",
            new StorageQueueInner(), com.azure.core.util.Context.NONE);
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
