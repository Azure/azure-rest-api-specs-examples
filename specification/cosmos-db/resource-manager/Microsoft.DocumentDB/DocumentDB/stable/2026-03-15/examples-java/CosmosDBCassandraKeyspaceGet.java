
/**
 * Samples for CassandraResources GetCassandraKeyspace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBCassandraKeyspaceGet.json
     */
    /**
     * Sample code: CosmosDBCassandraKeyspaceGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBCassandraKeyspaceGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCassandraResources().getCassandraKeyspaceWithResponse("rg1", "ddb1", "keyspaceName",
            com.azure.core.util.Context.NONE);
    }
}
