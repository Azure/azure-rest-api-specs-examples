
/**
 * Samples for CassandraResources GetCassandraTableThroughput.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBCassandraTableThroughputGet.json
     */
    /**
     * Sample code: CosmosDBCassandraTableThroughputGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBCassandraTableThroughputGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCassandraResources().getCassandraTableThroughputWithResponse("rg1", "ddb1",
            "keyspaceName", "tableName", com.azure.core.util.Context.NONE);
    }
}
