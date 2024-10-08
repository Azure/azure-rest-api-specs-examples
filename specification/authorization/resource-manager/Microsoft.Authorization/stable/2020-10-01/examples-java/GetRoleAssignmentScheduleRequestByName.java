
/**
 * Samples for RoleAssignmentScheduleRequests Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/
     * GetRoleAssignmentScheduleRequestByName.json
     */
    /**
     * Sample code: GetRoleAssignmentScheduleRequestByName.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getRoleAssignmentScheduleRequestByName(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.accessManagement().roleAssignments().manager().roleServiceClient().getRoleAssignmentScheduleRequests()
            .getWithResponse("providers/Microsoft.Subscription/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f",
                "fea7a502-9a96-4806-a26f-eee560e52045", com.azure.core.util.Context.NONE);
    }
}
