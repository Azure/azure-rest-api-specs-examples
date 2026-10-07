
/**
 * Samples for GraphResources ListGraphs.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBGraphResourceList.json
     */
    /**
     * Sample code: CosmosDBSqlDatabaseList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSqlDatabaseList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getGraphResources().listGraphs("rgName", "ddb1", com.azure.core.util.Context.NONE);
    }
}
