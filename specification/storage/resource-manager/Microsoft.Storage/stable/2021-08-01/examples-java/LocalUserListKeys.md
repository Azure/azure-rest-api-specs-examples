```java
import com.azure.core.util.Context;

/** Samples for LocalUsersOperation ListKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/LocalUserListKeys.json
     */
    /**
     * Sample code: ListLocalUserKeys.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listLocalUserKeys(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getLocalUsersOperations()
            .listKeysWithResponse("res6977", "sto2527", "user1", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
