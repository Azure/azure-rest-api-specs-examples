
/**
 * Samples for RoleAssignmentSchedules Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/
     * GetRoleAssignmentScheduleByName.json
     */
    /**
     * Sample code: GetRoleAssignmentScheduleByName.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getRoleAssignmentScheduleByName(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.accessManagement().roleAssignments().manager().roleServiceClient().getRoleAssignmentSchedules()
            .getWithResponse("providers/Microsoft.Subscription/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f",
                "c9e264ff-3133-4776-a81a-ebc7c33c8ec6", com.azure.core.util.Context.NONE);
    }
}
