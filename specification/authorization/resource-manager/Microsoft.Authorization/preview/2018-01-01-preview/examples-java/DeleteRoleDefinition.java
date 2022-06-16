import com.azure.core.util.Context;

/** Samples for RoleDefinitions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/preview/2018-01-01-preview/examples/DeleteRoleDefinition.json
     */
    /**
     * Sample code: Delete role definition.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteRoleDefinition(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .accessManagement()
            .roleAssignments()
            .manager()
            .roleServiceClient()
            .getRoleDefinitions()
            .deleteWithResponse("scope", "roleDefinitionId", Context.NONE);
    }
}
