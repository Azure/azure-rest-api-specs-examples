import com.azure.core.util.Context;

/** Samples for MongoDBResources GetMongoDBCollectionThroughput. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-08-15/examples/CosmosDBMongoDBCollectionThroughputGet.json
     */
    /**
     * Sample code: CosmosDBMongoDBCollectionThroughputGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBMongoDBCollectionThroughputGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getMongoDBResources()
            .getMongoDBCollectionThroughputWithResponse("rg1", "ddb1", "databaseName", "collectionName", Context.NONE);
    }
}
