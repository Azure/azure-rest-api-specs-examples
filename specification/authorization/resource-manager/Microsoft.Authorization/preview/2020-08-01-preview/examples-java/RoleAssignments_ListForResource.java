import com.azure.core.util.Context;

/** Samples for RoleAssignments ListForResource. */
public final class Main {
    /*
     * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/preview/2020-08-01-preview/examples/RoleAssignments_ListForResource.json
     */
    /**
     * Sample code: List role assignments for a resource.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listRoleAssignmentsForAResource(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .accessManagement()
            .roleAssignments()
            .manager()
            .roleServiceClient()
            .getRoleAssignments()
            .listForResource(
                "testrg", "Microsoft.DocumentDb", "databaseAccounts", "test-db-account", null, null, Context.NONE);
    }
}
