
/** Samples for DenyAssignments ListForScope. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/examples/
     * GetDenyAssignmentByScope.json
     */
    /**
     * Sample code: List deny assignments for scope.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listDenyAssignmentsForScope(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.accessManagement().roleAssignments().manager().roleServiceClient().getDenyAssignments()
            .listForScope("subscriptions/subId", null, com.azure.core.util.Context.NONE);
    }
}
