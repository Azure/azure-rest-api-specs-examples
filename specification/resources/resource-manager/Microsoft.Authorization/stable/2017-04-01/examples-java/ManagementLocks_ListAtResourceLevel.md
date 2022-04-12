Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ManagementLocks ListAtResourceLevel. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2017-04-01/examples/ManagementLocks_ListAtResourceLevel.json
     */
    /**
     * Sample code: List management locks at resource level.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listManagementLocksAtResourceLevel(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .genericResources()
            .manager()
            .managementLockClient()
            .getManagementLocks()
            .listAtResourceLevel(
                "resourcegroupname",
                "Microsoft.Storage",
                "parentResourcePath",
                "storageAccounts",
                "teststorageaccount",
                null,
                Context.NONE);
    }
}
```
