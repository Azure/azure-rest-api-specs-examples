Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Queue Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/QueueOperationDelete.json
     */
    /**
     * Sample code: QueueOperationDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void queueOperationDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getQueues()
            .deleteWithResponse("res3376", "sto328", "queue6185", Context.NONE);
    }
}
```
