
/**
 * Samples for CassandraResources GetCassandraViewThroughput.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBCassandraViewThroughputGet.json
     */
    /**
     * Sample code: CosmosDBCassandraViewThroughputGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBCassandraViewThroughputGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCassandraResources().getCassandraViewThroughputWithResponse("rg1", "ddb1",
            "keyspacename", "viewname", com.azure.core.util.Context.NONE);
    }
}
