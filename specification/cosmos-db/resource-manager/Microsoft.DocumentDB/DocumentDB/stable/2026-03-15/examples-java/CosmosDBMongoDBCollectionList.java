
/**
 * Samples for MongoDBResources ListMongoDBCollections.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBMongoDBCollectionList.json
     */
    /**
     * Sample code: CosmosDBMongoDBCollectionList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBMongoDBCollectionList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getMongoDBResources().listMongoDBCollections("rgName", "ddb1", "databaseName",
            com.azure.core.util.Context.NONE);
    }
}
