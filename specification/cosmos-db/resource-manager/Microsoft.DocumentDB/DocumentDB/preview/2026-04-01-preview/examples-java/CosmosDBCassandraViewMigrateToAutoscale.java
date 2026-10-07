
/**
 * Samples for CassandraResources MigrateCassandraViewToAutoscale.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBCassandraViewMigrateToAutoscale.json
     */
    /**
     * Sample code: CosmosDBCassandraViewMigrateToAutoscale.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBCassandraViewMigrateToAutoscale(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCassandraResources().migrateCassandraViewToAutoscale("rg1", "ddb1", "keyspacename",
            "viewname", com.azure.core.util.Context.NONE);
    }
}
