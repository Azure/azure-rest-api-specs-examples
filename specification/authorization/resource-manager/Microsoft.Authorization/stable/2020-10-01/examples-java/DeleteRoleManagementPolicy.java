
/**
 * Samples for RoleManagementPolicies Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/
     * DeleteRoleManagementPolicy.json
     */
    /**
     * Sample code: DeleteRoleManagementPolicy.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteRoleManagementPolicy(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.accessManagement().roleAssignments().manager().roleServiceClient().getRoleManagementPolicies()
            .deleteWithResponse("providers/Microsoft.Subscription/subscriptions/129ff972-28f8-46b8-a726-e497be039368",
                "570c3619-7688-4b34-b290-2b8bb3ccab2a", com.azure.core.util.Context.NONE);
    }
}
