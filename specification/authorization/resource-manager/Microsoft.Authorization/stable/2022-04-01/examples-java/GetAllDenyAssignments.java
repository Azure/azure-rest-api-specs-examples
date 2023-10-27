/** Samples for DenyAssignments List. */
public final class Main {
    /*
     * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/examples/GetAllDenyAssignments.json
     */
    /**
     * Sample code: List deny assignments for subscription.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listDenyAssignmentsForSubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .accessManagement()
            .roleAssignments()
            .manager()
            .roleServiceClient()
            .getDenyAssignments()
            .list(null, com.azure.core.util.Context.NONE);
    }
}
