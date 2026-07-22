
/**
 * Samples for MongoDBResources GetMongoDBCollectionThroughput.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBMongoDBCollectionThroughputGet.json
     */
    /**
     * Sample code: CosmosDBMongoDBCollectionThroughputGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBMongoDBCollectionThroughputGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getMongoDBResources().getMongoDBCollectionThroughputWithResponse("rg1", "ddb1",
            "databaseName", "collectionName", com.azure.core.util.Context.NONE);
    }
}
