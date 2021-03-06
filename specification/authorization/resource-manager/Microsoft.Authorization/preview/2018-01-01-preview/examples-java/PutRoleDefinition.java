import com.azure.core.util.Context;
import com.azure.resourcemanager.authorization.fluent.models.RoleDefinitionInner;

/** Samples for RoleDefinitions CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/preview/2018-01-01-preview/examples/PutRoleDefinition.json
     */
    /**
     * Sample code: Create role definition.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createRoleDefinition(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .accessManagement()
            .roleAssignments()
            .manager()
            .roleServiceClient()
            .getRoleDefinitions()
            .createOrUpdateWithResponse("scope", "roleDefinitionId", new RoleDefinitionInner(), Context.NONE);
    }
}
