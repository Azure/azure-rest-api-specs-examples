
/**
 * Samples for CassandraResources DeleteCassandraKeyspace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBCassandraKeyspaceDelete.json
     */
    /**
     * Sample code: CosmosDBCassandraKeyspaceDelete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBCassandraKeyspaceDelete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCassandraResources().deleteCassandraKeyspace("rg1", "ddb1", "keyspaceName",
            com.azure.core.util.Context.NONE);
    }
}
