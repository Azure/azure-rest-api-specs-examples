import com.azure.core.util.Context;

/** Samples for DenyAssignments Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/preview/2018-07-01-preview/examples/GetDenyAssignmentByNameId.json
     */
    /**
     * Sample code: Get deny assignment by name.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getDenyAssignmentByName(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .accessManagement()
            .roleAssignments()
            .manager()
            .roleServiceClient()
            .getDenyAssignments()
            .getWithResponse("subscriptions/subId/resourcegroups/rgname", "denyAssignmentId", Context.NONE);
    }
}
