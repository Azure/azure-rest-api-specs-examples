import com.azure.core.util.Context;

/** Samples for MongoDBResources ListMongoDBCollections. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBMongoDBCollectionList.json
     */
    /**
     * Sample code: CosmosDBMongoDBCollectionList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBMongoDBCollectionList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getMongoDBResources()
            .listMongoDBCollections("rgName", "ddb1", "databaseName", Context.NONE);
    }
}
