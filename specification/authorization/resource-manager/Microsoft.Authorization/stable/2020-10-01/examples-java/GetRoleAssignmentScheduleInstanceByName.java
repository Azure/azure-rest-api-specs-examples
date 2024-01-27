
/** Samples for RoleAssignmentScheduleInstances Get. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/
     * GetRoleAssignmentScheduleInstanceByName.json
     */
    /**
     * Sample code: GetRoleAssignmentScheduleInstanceByName.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getRoleAssignmentScheduleInstanceByName(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.accessManagement().roleAssignments().manager().roleServiceClient().getRoleAssignmentScheduleInstances()
            .getWithResponse("providers/Microsoft.Subscription/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f",
                "ed9b8180-cef7-4c77-a63c-b8566ecfc412", com.azure.core.util.Context.NONE);
    }
}
