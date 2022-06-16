import com.azure.core.util.Context;

/** Samples for RoleDefinitions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/preview/2018-01-01-preview/examples/GetRoleDefinitionAtScope.json
     */
    /**
     * Sample code: List role definition for scope.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listRoleDefinitionForScope(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .accessManagement()
            .roleAssignments()
            .manager()
            .roleServiceClient()
            .getRoleDefinitions()
            .list("scope", null, Context.NONE);
    }
}
