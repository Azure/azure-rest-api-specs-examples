import com.azure.core.util.Context;

/** Samples for RoleManagementPolicyAssignments Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/GetRoleManagementPolicyAssignmentByName.json
     */
    /**
     * Sample code: GetConfigurations.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getConfigurations(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .accessManagement()
            .roleAssignments()
            .manager()
            .roleServiceClient()
            .getRoleManagementPolicyAssignments()
            .getWithResponse(
                "providers/Microsoft.Subscription/subscriptions/129ff972-28f8-46b8-a726-e497be039368",
                "b959d571-f0b5-4042-88a7-01be6cb22db9_a1705bd2-3a8f-45a5-8683-466fcfd5cc24",
                Context.NONE);
    }
}
