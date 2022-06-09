```java
import com.azure.core.util.Context;

/** Samples for RoleManagementPolicies ListForScope. */
public final class Main {
    /*
     * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/GetRoleManagementPolicyByScope.json
     */
    /**
     * Sample code: GetRoleManagementPolicyByRoleDefinitionFilter.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getRoleManagementPolicyByRoleDefinitionFilter(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .accessManagement()
            .roleAssignments()
            .manager()
            .roleServiceClient()
            .getRoleManagementPolicies()
            .listForScope(
                "providers/Microsoft.Subscription/subscriptions/129ff972-28f8-46b8-a726-e497be039368", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
