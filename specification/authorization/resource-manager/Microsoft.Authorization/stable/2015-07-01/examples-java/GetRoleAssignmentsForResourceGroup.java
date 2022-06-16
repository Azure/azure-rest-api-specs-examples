import com.azure.core.util.Context;

/** Samples for RoleAssignments ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/stable/2015-07-01/examples/GetRoleAssignmentsForResourceGroup.json
     */
    /**
     * Sample code: List role assignments for resource group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listRoleAssignmentsForResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .accessManagement()
            .roleAssignments()
            .manager()
            .roleServiceClient()
            .getRoleAssignments()
            .listByResourceGroup("rgname", null, Context.NONE);
    }
}
