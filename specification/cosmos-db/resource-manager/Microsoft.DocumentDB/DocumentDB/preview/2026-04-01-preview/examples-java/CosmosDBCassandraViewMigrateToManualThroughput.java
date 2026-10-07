
/**
 * Samples for CassandraResources MigrateCassandraViewToManualThroughput.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBCassandraViewMigrateToManualThroughput.json
     */
    /**
     * Sample code: CosmosDBCassandraViewMigrateToManualThroughput.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        cosmosDBCassandraViewMigrateToManualThroughput(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCassandraResources().migrateCassandraViewToManualThroughput("rg1", "ddb1",
            "keyspacename", "viewname", com.azure.core.util.Context.NONE);
    }
}
