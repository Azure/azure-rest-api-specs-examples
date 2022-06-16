import com.azure.core.util.Context;

/** Samples for RoleAssignments DeleteById. */
public final class Main {
    /*
     * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/stable/2015-07-01/examples/DeleteRoleAssignmentById.json
     */
    /**
     * Sample code: Delete role assignment by ID.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteRoleAssignmentByID(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .accessManagement()
            .roleAssignments()
            .manager()
            .roleServiceClient()
            .getRoleAssignments()
            .deleteByIdWithResponse("roleAssignmentId", Context.NONE);
    }
}
