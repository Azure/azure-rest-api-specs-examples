
/**
 * Samples for MongoDBResources GetMongoDBDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-11-15/examples/
     * CosmosDBMongoDBDatabaseGet.json
     */
    /**
     * Sample code: CosmosDBMongoDBDatabaseGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBMongoDBDatabaseGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getMongoDBResources().getMongoDBDatabaseWithResponse("rg1",
            "ddb1", "databaseName", com.azure.core.util.Context.NONE);
    }
}
