```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.storage.models.StorageAccountRegenerateKeyParameters;

/** Samples for StorageAccounts RegenerateKey. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/StorageAccountRegenerateKey.json
     */
    /**
     * Sample code: StorageAccountRegenerateKey.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountRegenerateKey(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getStorageAccounts()
            .regenerateKeyWithResponse(
                "res4167", "sto3539", new StorageAccountRegenerateKeyParameters().withKeyName("key2"), Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
