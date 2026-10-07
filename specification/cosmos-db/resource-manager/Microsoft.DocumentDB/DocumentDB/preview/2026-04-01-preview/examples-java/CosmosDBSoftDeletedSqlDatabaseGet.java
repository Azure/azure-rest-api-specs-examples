
/**
 * Samples for SoftDeletedSqlDatabases Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBSoftDeletedSqlDatabaseGet.json
     */
    /**
     * Sample code: CosmosDBSoftDeletedSqlDatabaseGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSoftDeletedSqlDatabaseGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSoftDeletedSqlDatabases().getWithResponse("rg1", "West US", "softdeleted-cosmosdb-1",
            "MyDatabase", com.azure.core.util.Context.NONE);
    }
}
