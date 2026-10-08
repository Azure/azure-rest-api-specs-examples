
/**
 * Samples for SoftDeletedSqlDatabases Purge.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBSoftDeletedSqlDatabasePurge.json
     */
    /**
     * Sample code: CosmosDBSoftDeletedSqlDatabasePurge.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSoftDeletedSqlDatabasePurge(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSoftDeletedSqlDatabases().purge("rg1", "West US", "softdeleted-cosmosdb-1",
            "MyDatabase", null, com.azure.core.util.Context.NONE);
    }
}
