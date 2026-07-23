
/**
 * Samples for CassandraResources MigrateCassandraTableToAutoscale.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBCassandraTableMigrateToAutoscale.json
     */
    /**
     * Sample code: CosmosDBCassandraTableMigrateToAutoscale.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        cosmosDBCassandraTableMigrateToAutoscale(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCassandraResources().migrateCassandraTableToAutoscale("rg1", "ddb1", "keyspaceName",
            "tableName", com.azure.core.util.Context.NONE);
    }
}
