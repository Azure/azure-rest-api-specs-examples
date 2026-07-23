
/**
 * Samples for SqlResources ListSqlDatabases.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBSqlDatabaseList.json
     */
    /**
     * Sample code: CosmosDBSqlDatabaseList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSqlDatabaseList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSqlResources().listSqlDatabases("rgName", "ddb1", com.azure.core.util.Context.NONE);
    }
}
