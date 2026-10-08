
/**
 * Samples for CassandraResources GetCassandraView.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBCassandraViewGet.json
     */
    /**
     * Sample code: CosmosDBCassandraViewGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBCassandraViewGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCassandraResources().getCassandraViewWithResponse("rg1", "ddb1", "keyspacename",
            "viewname", com.azure.core.util.Context.NONE);
    }
}
