
/**
 * Samples for RoleEligibilityScheduleRequests ListForScope.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/
     * GetRoleEligibilityScheduleRequestByScope.json
     */
    /**
     * Sample code: GetRoleEligibilityScheduleRequestByScope.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getRoleEligibilityScheduleRequestByScope(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.accessManagement().roleAssignments().manager().roleServiceClient().getRoleEligibilityScheduleRequests()
            .listForScope("providers/Microsoft.Subscription/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f",
                "assignedTo('A3BB8764-CB92-4276-9D2A-CA1E895E55EA')", com.azure.core.util.Context.NONE);
    }
}
