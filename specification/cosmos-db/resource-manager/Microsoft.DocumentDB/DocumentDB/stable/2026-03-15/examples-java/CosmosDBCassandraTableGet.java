
/**
 * Samples for CassandraResources GetCassandraTable.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBCassandraTableGet.json
     */
    /**
     * Sample code: CosmosDBCassandraTableGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBCassandraTableGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCassandraResources().getCassandraTableWithResponse("rg1", "ddb1", "keyspaceName",
            "tableName", com.azure.core.util.Context.NONE);
    }
}
