import com.azure.core.util.Context;

/** Samples for RoleManagementPolicyAssignments ListForScope. */
public final class Main {
    /*
     * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/GetRoleManagementPolicyAssignmentByScope.json
     */
    /**
     * Sample code: GetRoleManagementPolicyAssignmentByScope.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getRoleManagementPolicyAssignmentByScope(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .accessManagement()
            .roleAssignments()
            .manager()
            .roleServiceClient()
            .getRoleManagementPolicyAssignments()
            .listForScope(
                "providers/Microsoft.Subscription/subscriptions/129ff972-28f8-46b8-a726-e497be039368", Context.NONE);
    }
}
