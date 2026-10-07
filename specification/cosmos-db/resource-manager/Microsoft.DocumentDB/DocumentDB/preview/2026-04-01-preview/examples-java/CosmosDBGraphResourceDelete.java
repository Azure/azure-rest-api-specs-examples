
/**
 * Samples for GraphResources DeleteGraphResource.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBGraphResourceDelete.json
     */
    /**
     * Sample code: CosmosDBSqlDatabaseDelete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSqlDatabaseDelete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getGraphResources().deleteGraphResource("rg1", "ddb1", "graphName",
            com.azure.core.util.Context.NONE);
    }
}
