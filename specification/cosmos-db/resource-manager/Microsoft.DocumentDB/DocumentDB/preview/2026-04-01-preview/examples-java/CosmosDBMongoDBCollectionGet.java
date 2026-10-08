
/**
 * Samples for MongoDBResources GetMongoDBCollection.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBMongoDBCollectionGet.json
     */
    /**
     * Sample code: CosmosDBMongoDBCollectionGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBMongoDBCollectionGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getMongoDBResources().getMongoDBCollectionWithResponse("rgName", "ddb1", "databaseName",
            "collectionName", com.azure.core.util.Context.NONE);
    }
}
