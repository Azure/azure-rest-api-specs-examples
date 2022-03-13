Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.storage.fluent.models.LocalUserInner;

/** Samples for LocalUsersOperation CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/LocalUserUpdate.json
     */
    /**
     * Sample code: UpdateLocalUser.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateLocalUser(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getLocalUsersOperations()
            .createOrUpdateWithResponse(
                "res6977",
                "sto2527",
                "user1",
                new LocalUserInner()
                    .withHomeDirectory("homedirectory2")
                    .withHasSharedKey(false)
                    .withHasSshKey(false)
                    .withHasSshPassword(false),
                Context.NONE);
    }
}
```
