
/**
 * Samples for RestorableSqlDatabases List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBRestorableSqlDatabaseList.json
     */
    /**
     * Sample code: CosmosDBRestorableSqlDatabaseList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBRestorableSqlDatabaseList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getRestorableSqlDatabases().list("WestUS", "d9b26648-2f53-4541-b3d8-3044f4f9810d",
            com.azure.core.util.Context.NONE);
    }
}
