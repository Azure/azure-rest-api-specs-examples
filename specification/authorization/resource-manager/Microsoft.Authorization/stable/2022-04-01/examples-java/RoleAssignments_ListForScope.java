
/**
 * Samples for RoleAssignments ListForScope.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/examples/
     * RoleAssignments_ListForScope.json
     */
    /**
     * Sample code: List role assignments for scope.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listRoleAssignmentsForScope(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.accessManagement().roleAssignments().manager().roleServiceClient().getRoleAssignments().listForScope(
            "subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2", null, null, null, com.azure.core.util.Context.NONE);
    }
}
