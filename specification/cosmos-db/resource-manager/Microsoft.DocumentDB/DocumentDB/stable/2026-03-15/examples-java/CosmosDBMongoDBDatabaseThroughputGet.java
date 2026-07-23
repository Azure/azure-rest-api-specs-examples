
/**
 * Samples for MongoDBResources GetMongoDBDatabaseThroughput.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBMongoDBDatabaseThroughputGet.json
     */
    /**
     * Sample code: CosmosDBMongoDBDatabaseThroughputGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBMongoDBDatabaseThroughputGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getMongoDBResources().getMongoDBDatabaseThroughputWithResponse("rg1", "ddb1",
            "databaseName", com.azure.core.util.Context.NONE);
    }
}
