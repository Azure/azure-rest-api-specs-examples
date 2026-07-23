
/**
 * Samples for GremlinResources DeleteGremlinRoleDefinition.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/gremlinrbac/CosmosDBGremlinRoleDefinitionDelete.json
     */
    /**
     * Sample code: CosmosDBGremlinRoleDefinitionDelete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBGremlinRoleDefinitionDelete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getGremlinResources().deleteGremlinRoleDefinition("myResourceGroupName",
            "myAccountName", "myRoleDefinitionId", com.azure.core.util.Context.NONE);
    }
}
