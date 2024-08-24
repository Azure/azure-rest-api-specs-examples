
import com.azure.resourcemanager.authorization.fluent.models.RoleDefinitionInner;

/**
 * Samples for RoleDefinitions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/examples/PutRoleDefinition
     * .json
     */
    /**
     * Sample code: Create role definition.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createRoleDefinition(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.accessManagement().roleAssignments().manager().roleServiceClient().getRoleDefinitions()
            .createOrUpdateWithResponse("scope", "roleDefinitionId", new RoleDefinitionInner(),
                com.azure.core.util.Context.NONE);
    }
}
