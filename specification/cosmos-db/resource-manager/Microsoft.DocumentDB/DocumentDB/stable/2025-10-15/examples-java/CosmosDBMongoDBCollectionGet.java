
/**
 * Samples for MongoDBResources GetMongoDBCollection.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBMongoDBCollectionGet.json
     */
    /**
     * Sample code: CosmosDBMongoDBCollectionGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBMongoDBCollectionGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getMongoDBResources().getMongoDBCollectionWithResponse(
            "rgName", "ddb1", "databaseName", "collectionName", com.azure.core.util.Context.NONE);
    }
}
