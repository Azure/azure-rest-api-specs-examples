```java
import com.azure.core.util.Context;

/** Samples for StorageAccounts ListKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/StorageAccountListKeys.json
     */
    /**
     * Sample code: StorageAccountListKeys.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountListKeys(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getStorageAccounts()
            .listKeysWithResponse("res418", "sto2220", null, Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
