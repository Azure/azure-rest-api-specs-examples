
/**
 * Samples for GremlinResources GetGremlinRoleDefinition.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/gremlinrbac/CosmosDBGremlinRoleDefinitionGet.json
     */
    /**
     * Sample code: CosmosDBGremlinRoleDefinitionGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBGremlinRoleDefinitionGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getGremlinResources().getGremlinRoleDefinitionWithResponse("myResourceGroupName",
            "myAccountName", "myRoleDefinitionId", com.azure.core.util.Context.NONE);
    }
}
