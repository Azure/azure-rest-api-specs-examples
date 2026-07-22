
/**
 * Samples for CassandraResources MigrateCassandraKeyspaceToAutoscale.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBCassandraKeyspaceMigrateToAutoscale.json
     */
    /**
     * Sample code: CosmosDBCassandraKeyspaceMigrateToAutoscale.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        cosmosDBCassandraKeyspaceMigrateToAutoscale(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCassandraResources().migrateCassandraKeyspaceToAutoscale("rg1", "ddb1",
            "keyspaceName", com.azure.core.util.Context.NONE);
    }
}
