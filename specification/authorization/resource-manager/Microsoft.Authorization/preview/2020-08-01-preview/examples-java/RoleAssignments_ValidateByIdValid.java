import com.azure.core.util.Context;
import com.azure.resourcemanager.authorization.models.PrincipalType;
import com.azure.resourcemanager.authorization.models.RoleAssignmentCreateParameters;

/** Samples for RoleAssignments ValidateById. */
public final class Main {
    /*
     * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/preview/2020-08-01-preview/examples/RoleAssignments_ValidateByIdValid.json
     */
    /**
     * Sample code: Validate a role assignment create or update operation by ID with successful validation.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void validateARoleAssignmentCreateOrUpdateOperationByIDWithSuccessfulValidation(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .accessManagement()
            .roleAssignments()
            .manager()
            .roleServiceClient()
            .getRoleAssignments()
            .validateByIdWithResponse(
                "subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/providers/Microsoft.Authorization/roleAssignments/b0f43c54-e787-4862-89b1-a653fa9cf747",
                new RoleAssignmentCreateParameters()
                    .withRoleDefinitionId(
                        "/providers/Microsoft.Authorization/roleDefinitions/0b5fe924-9a61-425c-96af-cfe6e287ca2d")
                    .withPrincipalId("ce2ce14e-85d7-4629-bdbc-454d0519d987")
                    .withPrincipalType(PrincipalType.USER),
                Context.NONE);
    }
}
