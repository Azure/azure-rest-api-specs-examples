import com.azure.core.util.Context;

/** Samples for MongoDBResources GetMongoDBDatabaseThroughput. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBMongoDBDatabaseThroughputGet.json
     */
    /**
     * Sample code: CosmosDBMongoDBDatabaseThroughputGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBMongoDBDatabaseThroughputGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getMongoDBResources()
            .getMongoDBDatabaseThroughputWithResponse("rg1", "ddb1", "databaseName", Context.NONE);
    }
}
