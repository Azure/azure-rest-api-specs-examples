Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.12.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for LocalUsersOperation RegeneratePassword. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/LocalUserRegeneratePassword.json
     */
    /**
     * Sample code: RegenerateLocalUserPassword.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void regenerateLocalUserPassword(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getLocalUsersOperations()
            .regeneratePasswordWithResponse("res6977", "sto2527", "user1", Context.NONE);
    }
}
```
