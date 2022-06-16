import com.azure.core.util.Context;
import com.azure.resourcemanager.authorization.models.RoleAssignmentCreateParameters;
import com.azure.resourcemanager.authorization.models.RoleAssignmentProperties;

/** Samples for RoleAssignments CreateById. */
public final class Main {
    /*
     * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/stable/2015-07-01/examples/PutRoleAssignmentById.json
     */
    /**
     * Sample code: Create role assignment by ID.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createRoleAssignmentByID(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .accessManagement()
            .roleAssignments()
            .manager()
            .roleServiceClient()
            .getRoleAssignments()
            .createByIdWithResponse(
                "roleAssignmentId",
                new RoleAssignmentCreateParameters()
                    .withProperties(
                        new RoleAssignmentProperties()
                            .withRoleDefinitionId(
                                "/subscriptions/4004a9fd-d58e-48dc-aeb2-4a4aec58606f/providers/Microsoft.Authorization/roleDefinitions/de139f84-1756-47ae-9be6-808fbbe84772")
                            .withPrincipalId("d93a38bc-d029-4160-bfb0-fbda779ac214")),
                Context.NONE);
    }
}
