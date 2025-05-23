
/**
 * Samples for RoleAssignments DeleteById.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/examples/
     * RoleAssignments_DeleteById.json
     */
    /**
     * Sample code: Delete role assignment by ID.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteRoleAssignmentByID(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.accessManagement().roleAssignments().manager().roleServiceClient().getRoleAssignments()
            .deleteByIdWithResponse(
                "subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/providers/Microsoft.Authorization/roleAssignments/b0f43c54-e787-4862-89b1-a653fa9cf747",
                null, com.azure.core.util.Context.NONE);
    }
}
