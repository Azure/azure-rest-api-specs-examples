
/**
 * Samples for CassandraResources DeleteCassandraTable.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBCassandraTableDelete.json
     */
    /**
     * Sample code: CosmosDBCassandraTableDelete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBCassandraTableDelete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCassandraResources().deleteCassandraTable("rg1", "ddb1", "keyspaceName", "tableName",
            com.azure.core.util.Context.NONE);
    }
}
