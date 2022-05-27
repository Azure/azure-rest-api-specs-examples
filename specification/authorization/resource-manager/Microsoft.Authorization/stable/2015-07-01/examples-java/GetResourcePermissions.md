Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Permissions ListForResource. */
public final class Main {
    /*
     * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/stable/2015-07-01/examples/GetResourcePermissions.json
     */
    /**
     * Sample code: List permissions for resource.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listPermissionsForResource(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .accessManagement()
            .roleAssignments()
            .manager()
            .roleServiceClient()
            .getPermissions()
            .listForResource(
                "rgname", "rpnamespace", "parentResourcePath", "resourceType", "resourceName", Context.NONE);
    }
}
```
