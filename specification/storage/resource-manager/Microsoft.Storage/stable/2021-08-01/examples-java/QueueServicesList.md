Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.12.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for QueueServices List. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/QueueServicesList.json
     */
    /**
     * Sample code: QueueServicesList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void queueServicesList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getQueueServices()
            .listWithResponse("res9290", "sto1590", Context.NONE);
    }
}
```
