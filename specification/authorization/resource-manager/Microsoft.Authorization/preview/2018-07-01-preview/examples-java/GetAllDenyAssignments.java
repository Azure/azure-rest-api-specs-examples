import com.azure.core.util.Context;

/** Samples for DenyAssignments List. */
public final class Main {
    /*
     * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/preview/2018-07-01-preview/examples/GetAllDenyAssignments.json
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
            .list(null, Context.NONE);
    }
}
