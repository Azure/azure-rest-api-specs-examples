
/**
 * Samples for CassandraResources MigrateCassandraTableToManualThroughput.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBCassandraTableMigrateToManualThroughput.json
     */
    /**
     * Sample code: CosmosDBCassandraTableMigrateToManualThroughput.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        cosmosDBCassandraTableMigrateToManualThroughput(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCassandraResources().migrateCassandraTableToManualThroughput("rg1", "ddb1",
            "keyspaceName", "tableName", com.azure.core.util.Context.NONE);
    }
}
