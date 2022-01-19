Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for StorageAccounts RevokeUserDelegationKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-04-01/examples/StorageAccountRevokeUserDelegationKeys.json
     */
    /**
     * Sample code: StorageAccountRevokeUserDelegationKeys.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountRevokeUserDelegationKeys(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getStorageAccounts()
            .revokeUserDelegationKeysWithResponse("res4167", "sto3539", Context.NONE);
    }
}
```
