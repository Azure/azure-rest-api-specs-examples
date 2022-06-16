import com.azure.core.util.Context;

/** Samples for RoleDefinitions GetById. */
public final class Main {
    /*
     * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/stable/2015-07-01/examples/GetRoleDefinitionById.json
     */
    /**
     * Sample code: Get role definition by ID.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getRoleDefinitionByID(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .accessManagement()
            .roleAssignments()
            .manager()
            .roleServiceClient()
            .getRoleDefinitions()
            .getByIdWithResponse("roleDefinitionId", Context.NONE);
    }
}
