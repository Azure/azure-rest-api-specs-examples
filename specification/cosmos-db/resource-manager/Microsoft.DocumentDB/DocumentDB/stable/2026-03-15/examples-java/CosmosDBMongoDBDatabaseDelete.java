
/**
 * Samples for MongoDBResources DeleteMongoDBDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBMongoDBDatabaseDelete.json
     */
    /**
     * Sample code: CosmosDBMongoDBDatabaseDelete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBMongoDBDatabaseDelete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getMongoDBResources().deleteMongoDBDatabase("rg1", "ddb1", "databaseName",
            com.azure.core.util.Context.NONE);
    }
}
