import com.azure.core.util.Context;

/** Samples for MongoDBResources DeleteMongoDBDatabase. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBMongoDBDatabaseDelete.json
     */
    /**
     * Sample code: CosmosDBMongoDBDatabaseDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBMongoDBDatabaseDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getMongoDBResources()
            .deleteMongoDBDatabase("rg1", "ddb1", "databaseName", Context.NONE);
    }
}
