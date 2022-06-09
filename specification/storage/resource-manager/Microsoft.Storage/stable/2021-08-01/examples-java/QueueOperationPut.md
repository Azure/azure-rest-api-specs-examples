```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.storage.fluent.models.StorageQueueInner;
import java.util.HashMap;
import java.util.Map;

/** Samples for Queue Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/QueueOperationPut.json
     */
    /**
     * Sample code: QueueOperationPut.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void queueOperationPut(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getQueues()
            .createWithResponse("res3376", "sto328", "queue6185", new StorageQueueInner(), Context.NONE);
    }

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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
