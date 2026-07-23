
/**
 * Samples for GremlinResources GetGremlinRoleAssignment.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/gremlinrbac/CosmosDBGremlinRoleAssignmentGet.json
     */
    /**
     * Sample code: CosmosDBGremlinRoleAssignmentGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBGremlinRoleAssignmentGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getGremlinResources().getGremlinRoleAssignmentWithResponse("myResourceGroupName",
            "myAccountName", "myRoleAssignmentId", com.azure.core.util.Context.NONE);
    }
}
