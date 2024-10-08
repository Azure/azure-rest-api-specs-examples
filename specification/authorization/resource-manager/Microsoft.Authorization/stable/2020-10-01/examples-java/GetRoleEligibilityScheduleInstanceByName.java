
/**
 * Samples for RoleEligibilityScheduleInstances Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/
     * GetRoleEligibilityScheduleInstanceByName.json
     */
    /**
     * Sample code: GetRoleEligibilityScheduleInstanceByName.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getRoleEligibilityScheduleInstanceByName(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.accessManagement().roleAssignments().manager().roleServiceClient().getRoleEligibilityScheduleInstances()
            .getWithResponse("providers/Microsoft.Subscription/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f",
                "21e4b59a-0499-4fe0-a3c3-43a3055b773a", com.azure.core.util.Context.NONE);
    }
}
