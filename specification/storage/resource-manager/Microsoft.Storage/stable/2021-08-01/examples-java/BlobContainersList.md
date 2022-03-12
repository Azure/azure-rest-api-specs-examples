Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for BlobContainers List. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/BlobContainersList.json
     */
    /**
     * Sample code: ListContainers.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listContainers(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getBlobContainers()
            .list("res9290", "sto1590", null, null, null, Context.NONE);
    }
}
```
