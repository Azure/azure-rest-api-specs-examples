
/**
 * Samples for CassandraResources ListCassandraViews.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBCassandraViewList.json
     */
    /**
     * Sample code: CosmosDBCassandraViewList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBCassandraViewList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCassandraResources().listCassandraViews("rgName", "ddb1", "keyspacename",
            com.azure.core.util.Context.NONE);
    }
}
