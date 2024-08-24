
/**
 * Samples for RoleEligibilityScheduleInstances ListForScope.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/
     * GetRoleEligibilityScheduleInstancesByScope.json
     */
    /**
     * Sample code: GetRoleEligibilityScheduleInstancesByScope.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getRoleEligibilityScheduleInstancesByScope(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.accessManagement().roleAssignments().manager().roleServiceClient().getRoleEligibilityScheduleInstances()
            .listForScope("providers/Microsoft.Subscription/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f",
                "assignedTo('a3bb8764-cb92-4276-9d2a-ca1e895e55ea')", com.azure.core.util.Context.NONE);
    }
}
