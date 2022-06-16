import com.azure.core.util.Context;

/** Samples for RoleAssignments Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/stable/2015-07-01/examples/DeleteRoleAssignmentByName.json
     */
    /**
     * Sample code: Delete role assignment by name.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteRoleAssignmentByName(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .accessManagement()
            .roleAssignments()
            .manager()
            .roleServiceClient()
            .getRoleAssignments()
            .deleteWithResponse("scope", "roleAssignmentName", Context.NONE);
    }
}
