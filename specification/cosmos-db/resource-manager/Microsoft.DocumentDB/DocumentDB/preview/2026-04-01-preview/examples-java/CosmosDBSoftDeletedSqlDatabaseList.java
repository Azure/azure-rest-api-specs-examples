
/**
 * Samples for SoftDeletedSqlDatabases List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBSoftDeletedSqlDatabaseList.json
     */
    /**
     * Sample code: CosmosDBSoftDeletedSqlDatabaseList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSoftDeletedSqlDatabaseList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSoftDeletedSqlDatabases().listWithResponse("rg1", "West US",
            "softdeleted-cosmosdb-1", com.azure.core.util.Context.NONE);
    }
}
