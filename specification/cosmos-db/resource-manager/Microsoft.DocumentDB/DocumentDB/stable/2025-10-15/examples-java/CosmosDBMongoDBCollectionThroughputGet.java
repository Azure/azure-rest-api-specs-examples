
/**
 * Samples for MongoDBResources GetMongoDBCollectionThroughput.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBMongoDBCollectionThroughputGet.json
     */
    /**
     * Sample code: CosmosDBMongoDBCollectionThroughputGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBMongoDBCollectionThroughputGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getMongoDBResources()
            .getMongoDBCollectionThroughputWithResponse("rg1", "ddb1", "databaseName", "collectionName",
                com.azure.core.util.Context.NONE);
    }
}
