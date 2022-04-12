Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for FileShares Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/FileSharesDelete.json
     */
    /**
     * Sample code: DeleteShares.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteShares(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getFileShares()
            .deleteWithResponse("res4079", "sto4506", "share9689", null, null, Context.NONE);
    }
}
```
