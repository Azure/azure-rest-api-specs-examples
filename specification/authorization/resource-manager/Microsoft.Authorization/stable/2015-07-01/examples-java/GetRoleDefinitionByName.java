import com.azure.core.util.Context;

/** Samples for RoleDefinitions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/stable/2015-07-01/examples/GetRoleDefinitionByName.json
     */
    /**
     * Sample code: Get role definition by name.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getRoleDefinitionByName(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .accessManagement()
            .roleAssignments()
            .manager()
            .roleServiceClient()
            .getRoleDefinitions()
            .getWithResponse("scope", "roleDefinitionId", Context.NONE);
    }
}
