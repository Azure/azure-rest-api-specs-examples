
/**
 * Samples for GremlinResources ListGremlinRoleDefinitions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/gremlinrbac/CosmosDBGremlinRoleDefinitionList.json
     */
    /**
     * Sample code: CosmosDBGremlinRoleDefinitionList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBGremlinRoleDefinitionList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getGremlinResources().listGremlinRoleDefinitions("myResourceGroupName", "myAccountName",
            com.azure.core.util.Context.NONE);
    }
}
