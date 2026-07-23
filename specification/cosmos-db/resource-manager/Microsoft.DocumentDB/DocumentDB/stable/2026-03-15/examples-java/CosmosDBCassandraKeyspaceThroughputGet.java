
/**
 * Samples for CassandraResources GetCassandraKeyspaceThroughput.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBCassandraKeyspaceThroughputGet.json
     */
    /**
     * Sample code: CosmosDBCassandraKeyspaceThroughputGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBCassandraKeyspaceThroughputGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCassandraResources().getCassandraKeyspaceThroughputWithResponse("rg1", "ddb1",
            "keyspaceName", com.azure.core.util.Context.NONE);
    }
}
