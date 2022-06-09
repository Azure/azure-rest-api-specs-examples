```java
import com.azure.core.util.Context;

/** Samples for LocalUsersOperation Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/LocalUserGet.json
     */
    /**
     * Sample code: GetLocalUser.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getLocalUser(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getLocalUsersOperations()
            .getWithResponse("res6977", "sto2527", "user1", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
