
/**
 * Samples for GremlinResources ListGremlinRoleAssignments.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/gremlinrbac/CosmosDBGremlinRoleAssignmentList.json
     */
    /**
     * Sample code: CosmosDBGremlinRoleAssignmentList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBGremlinRoleAssignmentList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getGremlinResources().listGremlinRoleAssignments("myResourceGroupName", "myAccountName",
            com.azure.core.util.Context.NONE);
    }
}
