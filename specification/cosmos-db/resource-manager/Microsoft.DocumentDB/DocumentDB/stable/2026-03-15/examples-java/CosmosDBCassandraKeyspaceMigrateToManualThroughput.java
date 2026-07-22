
/**
 * Samples for CassandraResources MigrateCassandraKeyspaceToManualThroughput.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBCassandraKeyspaceMigrateToManualThroughput.json
     */
    /**
     * Sample code: CosmosDBCassandraKeyspaceMigrateToManualThroughput.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        cosmosDBCassandraKeyspaceMigrateToManualThroughput(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCassandraResources().migrateCassandraKeyspaceToManualThroughput("rg1", "ddb1",
            "keyspaceName", com.azure.core.util.Context.NONE);
    }
}
