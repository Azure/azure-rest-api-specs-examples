
/**
 * Samples for RoleEligibilitySchedules Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/
     * GetRoleEligibilityScheduleByName.json
     */
    /**
     * Sample code: GetRoleEligibilityScheduleByName.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getRoleEligibilityScheduleByName(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.accessManagement().roleAssignments().manager().roleServiceClient().getRoleEligibilitySchedules()
            .getWithResponse("providers/Microsoft.Subscription/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f",
                "b1477448-2cc6-4ceb-93b4-54a202a89413", com.azure.core.util.Context.NONE);
    }
}
