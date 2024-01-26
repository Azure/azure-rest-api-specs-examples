
/** Samples for RoleAssignments List. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/examples/
     * RoleAssignments_ListForSubscription.json
     */
    /**
     * Sample code: List role assignments for subscription.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listRoleAssignmentsForSubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.accessManagement().roleAssignments().manager().roleServiceClient().getRoleAssignments().list(null, null,
            com.azure.core.util.Context.NONE);
    }
}
