
/**
 * Samples for GremlinResources DeleteGremlinRoleAssignment.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/gremlinrbac/CosmosDBGremlinRoleAssignmentDelete.json
     */
    /**
     * Sample code: CosmosDBGremlinRoleAssignmentDelete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBGremlinRoleAssignmentDelete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getGremlinResources().deleteGremlinRoleAssignment("myResourceGroupName",
            "myAccountName", "myRoleAssignmentId", com.azure.core.util.Context.NONE);
    }
}
