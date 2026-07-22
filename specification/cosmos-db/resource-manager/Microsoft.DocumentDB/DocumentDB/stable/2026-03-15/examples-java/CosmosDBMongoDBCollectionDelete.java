
/**
 * Samples for MongoDBResources DeleteMongoDBCollection.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBMongoDBCollectionDelete.json
     */
    /**
     * Sample code: CosmosDBMongoDBCollectionDelete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBMongoDBCollectionDelete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getMongoDBResources().deleteMongoDBCollection("rg1", "ddb1", "databaseName",
            "collectionName", com.azure.core.util.Context.NONE);
    }
}
