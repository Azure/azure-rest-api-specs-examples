Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for RoleAssignments ListForScope. */
public final class Main {
    /*
     * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/preview/2020-08-01-preview/examples/RoleAssignments_ListForScope.json
     */
    /**
     * Sample code: List role assignments for scope.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listRoleAssignmentsForScope(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .accessManagement()
            .roleAssignments()
            .manager()
            .roleServiceClient()
            .getRoleAssignments()
            .listForScope("subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2", null, null, Context.NONE);
    }
}
```
