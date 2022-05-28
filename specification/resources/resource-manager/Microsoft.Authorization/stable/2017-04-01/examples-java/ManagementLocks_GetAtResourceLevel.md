Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ManagementLocks GetAtResourceLevel. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2017-04-01/examples/ManagementLocks_GetAtResourceLevel.json
     */
    /**
     * Sample code: Get management lock at resource level.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getManagementLockAtResourceLevel(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .genericResources()
            .manager()
            .managementLockClient()
            .getManagementLocks()
            .getAtResourceLevelWithResponse(
                "resourcegroupname",
                "Microsoft.Storage",
                "parentResourcePath",
                "storageAccounts",
                "teststorageaccount",
                "testlock",
                Context.NONE);
    }
}
```
