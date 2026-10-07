
/**
 * Samples for SoftDeletedSqlDatabases Restore.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBSoftDeletedSqlDatabaseRestore.json
     */
    /**
     * Sample code: CosmosDBSoftDeletedSqlDatabaseRestore.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSoftDeletedSqlDatabaseRestore(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSoftDeletedSqlDatabases().restore("rg1", "West US", "ddb1", "softDeletedDatabase1",
            null, com.azure.core.util.Context.NONE);
    }
}
