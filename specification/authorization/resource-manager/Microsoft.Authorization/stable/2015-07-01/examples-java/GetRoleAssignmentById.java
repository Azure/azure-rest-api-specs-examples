import com.azure.core.util.Context;

/** Samples for RoleAssignments GetById. */
public final class Main {
    /*
     * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/stable/2015-07-01/examples/GetRoleAssignmentById.json
     */
    /**
     * Sample code: Get role assignment by ID.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getRoleAssignmentByID(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .accessManagement()
            .roleAssignments()
            .manager()
            .roleServiceClient()
            .getRoleAssignments()
            .getByIdWithResponse("roleassignmentId", Context.NONE);
    }
}
