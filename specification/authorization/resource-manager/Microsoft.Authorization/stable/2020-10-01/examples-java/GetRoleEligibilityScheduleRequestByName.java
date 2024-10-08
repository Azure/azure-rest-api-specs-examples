
/**
 * Samples for RoleEligibilityScheduleRequests Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/
     * GetRoleEligibilityScheduleRequestByName.json
     */
    /**
     * Sample code: GetRoleEligibilityScheduleRequestByName.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getRoleEligibilityScheduleRequestByName(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.accessManagement().roleAssignments().manager().roleServiceClient().getRoleEligibilityScheduleRequests()
            .getWithResponse("providers/Microsoft.Subscription/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f",
                "64caffb6-55c0-4deb-a585-68e948ea1ad6", com.azure.core.util.Context.NONE);
    }
}
