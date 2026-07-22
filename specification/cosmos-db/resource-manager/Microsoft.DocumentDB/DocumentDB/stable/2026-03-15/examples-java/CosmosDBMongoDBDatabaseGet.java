
/**
 * Samples for MongoDBResources GetMongoDBDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBMongoDBDatabaseGet.json
     */
    /**
     * Sample code: CosmosDBMongoDBDatabaseGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBMongoDBDatabaseGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getMongoDBResources().getMongoDBDatabaseWithResponse("rg1", "ddb1", "databaseName",
            com.azure.core.util.Context.NONE);
    }
}
