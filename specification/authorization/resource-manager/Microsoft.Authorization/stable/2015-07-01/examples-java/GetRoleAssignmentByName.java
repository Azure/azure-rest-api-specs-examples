import com.azure.core.util.Context;

/** Samples for RoleAssignments Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/stable/2015-07-01/examples/GetRoleAssignmentByName.json
     */
    /**
     * Sample code: Get role assignment by name.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getRoleAssignmentByName(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .accessManagement()
            .roleAssignments()
            .manager()
            .roleServiceClient()
            .getRoleAssignments()
            .getWithResponse("scope", "roleAssignmentName", Context.NONE);
    }
}
