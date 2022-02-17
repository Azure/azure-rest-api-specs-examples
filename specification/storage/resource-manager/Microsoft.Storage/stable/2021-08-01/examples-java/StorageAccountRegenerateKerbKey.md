Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.12.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.storage.models.StorageAccountRegenerateKeyParameters;

/** Samples for StorageAccounts RegenerateKey. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/StorageAccountRegenerateKerbKey.json
     */
    /**
     * Sample code: StorageAccountRegenerateKerbKey.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountRegenerateKerbKey(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getStorageAccounts()
            .regenerateKeyWithResponse(
                "res4167", "sto3539", new StorageAccountRegenerateKeyParameters().withKeyName("kerb1"), Context.NONE);
    }
}
```
