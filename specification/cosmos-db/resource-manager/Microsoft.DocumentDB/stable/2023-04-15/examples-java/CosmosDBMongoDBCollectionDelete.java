/** Samples for MongoDBResources DeleteMongoDBCollection. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-04-15/examples/CosmosDBMongoDBCollectionDelete.json
     */
    /**
     * Sample code: CosmosDBMongoDBCollectionDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBMongoDBCollectionDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getMongoDBResources()
            .deleteMongoDBCollection("rg1", "ddb1", "databaseName", "collectionName", com.azure.core.util.Context.NONE);
    }
}
