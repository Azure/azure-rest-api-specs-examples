```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.authorization.fluent.models.RoleManagementPolicyAssignmentInner;

/** Samples for RoleManagementPolicyAssignments Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/PutRoleManagementPolicyAssignment.json
     */
    /**
     * Sample code: PutRoleManagementPolicyAssignment.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void putRoleManagementPolicyAssignment(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .accessManagement()
            .roleAssignments()
            .manager()
            .roleServiceClient()
            .getRoleManagementPolicyAssignments()
            .createWithResponse(
                "providers/Microsoft.Subscription/subscriptions/129ff972-28f8-46b8-a726-e497be039368",
                "b959d571-f0b5-4042-88a7-01be6cb22db9_a1705bd2-3a8f-45a5-8683-466fcfd5cc24",
                new RoleManagementPolicyAssignmentInner()
                    .withScope("/subscriptions/129ff972-28f8-46b8-a726-e497be039368")
                    .withRoleDefinitionId(
                        "/subscriptions/129ff972-28f8-46b8-a726-e497be039368/providers/Microsoft.Authorization/roleDefinitions/a1705bd2-3a8f-45a5-8683-466fcfd5cc24")
                    .withPolicyId(
                        "/subscriptions/129ff972-28f8-46b8-a726-e497be039368/providers/Microsoft.Authorization/roleManagementPolicies/b959d571-f0b5-4042-88a7-01be6cb22db9"),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
